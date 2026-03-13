package models

import "time"

type OrderRow struct {
	ID              int       `json:"id" db:"id"`
	OrderID         int       `json:"order_id" db:"order_id"`
	ProductID       int       `json:"product_id" db:"product_id"`
	PurchaseAtPrice int       `json:"purchase_at_price" db:"purchase_at_price"`
	Quantity        int       `json:"quantity" db:"quantity"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
