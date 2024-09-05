CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255),
    order_id VARCHAR(255),
    total_amount NUMERIC(10, 2),
    payment_method VARCHAR(50),
    payment_details TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS refunds (
    id SERIAL PRIMARY KEY,
    transaction_id VARCHAR(255),
    refund_amount NUMERIC(10, 2),
    reason TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);