package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/viktare/go-commerce/internal/models"
)

func CreateProduct(pool *pgxpool.Pool, title string, code string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	query := `
		INSERT INTO products (title, code)
		VALUES ($1, $2)
		RETURNING id, title, code	 
	`
	var product models.Product

	err := pool.QueryRow(ctx, query, title, code).Scan(
		&product.ID,
		&product.TITLE,
		&product.CODE,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
