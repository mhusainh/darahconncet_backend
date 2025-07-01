BEGIN;

CREATE TABLE IF NOT EXISTS public.health_passports (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES public.users(id),
    passport_number VARCHAR(255) UNIQUE,
    expiry_date TIMESTAMPTZ,
    status VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

COMMIT;