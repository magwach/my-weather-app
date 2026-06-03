package services

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/magwach/my-weather-app/backend/internal/config"
	"github.com/magwach/my-weather-app/backend/internal/models"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrInvalidCreds = errors.New("invalid credentials")
	ErrUserNotFound = errors.New("user not found")
)

type AuthService struct {
	DB *pgxpool.Pool
}

func NewAuthService(db *pgxpool.Pool) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	query := `
		SELECT id, name, email, password_hash, created_at
		FROM users
		WHERE email = $1
	`

	err := s.DB.QueryRow(context.Background(), query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (s *AuthService) RegisterUser(req models.RegisterRequest) (*models.AuthResponse, error) {

	_, err := s.GetUserByEmail(req.Email)
	if err == nil {
		return nil, ErrUserExists
	}
	if err != nil && err != ErrUserNotFound {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, err
	}

	var userID string
	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err = s.DB.QueryRow(context.Background(),
		query,
		req.Name,
		req.Email,
		string(hashed),
	).Scan(&userID)

	if err != nil {
		return nil, err
	}

	token, err := s.generateJWT(userID)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		Name:  req.Name,
	}, nil
}

func (s *AuthService) LoginUser(req models.LoginRequest) (*models.AuthResponse, error) {

	user, err := s.GetUserByEmail(req.Email)
	if err != nil {
		return nil, ErrInvalidCreds
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, ErrInvalidCreds
	}

	token, err := s.generateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		Name:  user.Name,
	}, nil
}

func (s *AuthService) generateJWT(userID string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.JwtSecret)
}
