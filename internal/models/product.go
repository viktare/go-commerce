package models

import "time"

type Product struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Code      string    `json:"code" db:"code"`
	Price     int       `json:"price" db:"price"`
	Stock     int       `json:"stock" db:"stock"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
