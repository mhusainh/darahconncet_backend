BEGIN;

CREATE TABLE IF NOT EXISTS public.users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    role VARCHAR(255),
    password VARCHAR(255),
    gender VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(255),
    blood_type VARCHAR(255),
    birth_date TIMESTAMPTZ,
    address TEXT,
    reset_password_token VARCHAR(255),
    verify_email_token VARCHAR(255),
    is_verified INT DEFAULT 0,
    last_donation_date TIMESTAMPTZ,
    donation_count INT DEFAULT 0,
    public_id VARCHAR(255),
    url_file VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

COMMIT;