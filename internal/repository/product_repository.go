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
		&product.Title,
		&product.Code,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func GetProducts(pool *pgxpool.Pool) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	query := `
		SELECT id, title, code
		FROM products
	`

	products := make([]models.Product, 0)
	rows, err := pool.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Title, &p.Code)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	// check for errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil

}
