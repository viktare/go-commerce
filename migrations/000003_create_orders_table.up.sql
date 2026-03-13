CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    status VARCHAR(20) NOT NULL DEFAULT 'in_process',
    total BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);