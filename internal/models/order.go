package models

import "time"

type Order struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Status    string    `json:"status" db:"status"`
	Total     int64     `json:"total" db:"total"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
