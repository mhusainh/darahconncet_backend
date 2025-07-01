BEGIN;

CREATE TABLE IF NOT EXISTS public.hospitals (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    address TEXT,
    city VARCHAR(255),
    province VARCHAR(255),
    latitude FLOAT,
    longitude FLOAT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

COMMIT;