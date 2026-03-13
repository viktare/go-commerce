CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id),
    transaction_id VARCHAR NOT NULL UNIQUE,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    amount BIGINT NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'eur',
    created_at TIMESTAMP DEFAULT now()
);