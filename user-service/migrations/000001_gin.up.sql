CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS couriers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    phone_number VARCHAR(20),
    delivery_area TEXT,
    assigned_orders INTEGER[], 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);