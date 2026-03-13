CREATE TABLE IF NOT EXISTS order_rows (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id),
    product_id INT NOT NULL REFERENCES products(id),
    purchase_at_price INT NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);