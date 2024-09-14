create type orderStatus as enum (
    'NEW',
    'PROCESSING',
    'INVALID',
    'PROCESSED'
    );

CREATE TABLE IF NOT EXISTS orders(
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    number VARCHAR UNIQUE NOT NULL,
    status orderStatus DEFAULT 'NEW',
    accrual INTEGER,
    uploaded_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)

);
CREATE INDEX IF NOT EXISTS index_orders_on_user_id ON orders(user_id);