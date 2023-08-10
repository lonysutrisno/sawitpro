CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    phone_number VARCHAR(13) UNIQUE NOT NULL,
    full_name VARCHAR(60) NOT NULL,
    password_hash VARCHAR(128) NOT NULL,
    password_salt VARCHAR(32) NOT NULL,
    login_attempts INT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
