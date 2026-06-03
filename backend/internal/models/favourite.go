package models

import "time"

type Favourite struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}

type AddFavoriteRequest struct {
	City string `json:"city"`
}
