CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    home_address VARCHAR(50) NOT NULL,
    refresh_token TEXT,
    expired_at TIMESTAMP,
    password_hash TEXT
);


CREATE TABLE "courier" (
    id SERIAL PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    refresh_token TEXT,
    expired_at TIMESTAMP,
    password_hash TEXT
);

CREATE TABLE "restaurant" (
    id SERIAL PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    home_address VARCHAR(50) NOT NULL,
    refresh_token TEXT,
    expired_at TIMESTAMP,
    password_hash TEXT
);