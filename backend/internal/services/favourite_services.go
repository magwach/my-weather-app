package services

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/magwach/my-weather-app/backend/internal/models"
)

type FavoritesService struct {
	DB             *pgxpool.Pool
	WeatherService *WeatherService
}

func NewFavoritesService(db *pgxpool.Pool, ws *WeatherService) *FavoritesService {
	return &FavoritesService{
		DB:             db,
		WeatherService: ws,
	}
}

func (s *FavoritesService) AddFavorite(userID, city string) error {
	city = strings.ToLower(strings.TrimSpace(city))

	var count int

	err := s.DB.QueryRow(context.Background(),
		`SELECT COUNT(*) FROM favorites WHERE user_id = $1`,
		userID,
	).Scan(&count)

	if err != nil {
		return err
	}

	if count >= 3 {
		return errors.New("favorites limit reached, maximum 3 cities allowed")
	}

	_, err = s.WeatherService.GetCurrentWeather(city)
	if err != nil {
		if err.Error() == "city not found" {
			return errors.New("city not found")
		}
		return err
	}

	_, err = s.DB.Exec(context.Background(),
		`INSERT INTO favorites (user_id, city) VALUES ($1, $2)`,
		userID,
		city,
	)

	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			if pgErr.Code == pgerrcode.UniqueViolation {
				return errors.New("city already in favorites")
			}
		}

		return err
	}

	return nil
}

func (s *FavoritesService) GetFavorites(userID string) ([]models.FavoriteWithWeather, error) {

	rows, err := s.DB.Query(context.Background(),
		`SELECT city FROM favorites WHERE user_id = $1`,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []string

	for rows.Next() {
		var city string
		if err := rows.Scan(&city); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		items []models.FavoriteWithWeather
	)

	for _, city := range cities {

		wg.Add(1)

		go func(c string) {
			defer wg.Done()

			weather, err := s.WeatherService.GetCurrentWeather(c)
			if err != nil {
				return
			}

			mu.Lock()
			items = append(items, models.FavoriteWithWeather{
				City:    c,
				Weather: weather,
			})
			mu.Unlock()

		}(city)
	}

	wg.Wait()

	return items, nil
}

func (s *FavoritesService) RemoveFavorite(userID, city string) error {

	commandTag, err := s.DB.Exec(context.Background(),
		`DELETE FROM favorites WHERE user_id = $1 AND city = $2`,
		userID,
		city,
	)

	if err != nil {
		return err
	}

	rows := commandTag.RowsAffected()

	if rows == 0 {
		return errors.New("city not found in favorites")
	}

	return nil
}
