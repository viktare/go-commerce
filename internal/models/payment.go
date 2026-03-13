package models

import "time"

type Payment struct {
	ID            int       `json:"id" db:"id"`
	OrderID       int       `json:"order_id" db:"order_id"`
	TransactionID string    `json:"transaction_id" db:"transaction_id"`
	Status        string    `json:"status" db:"status"`
	Amount        int64     `json:"amount" db:"amount"`
	Currency      string    `json:"currency" db:"currency"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
