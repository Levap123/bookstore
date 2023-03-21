CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    book_id VARCHAR(255) NOT NULL,
    user_id INTEGER NOT NULL,
    added_at TIMESTAMP NOT NULL,
    status VARCHAR(255) NOT NULL
);
