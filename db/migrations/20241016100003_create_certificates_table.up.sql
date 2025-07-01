BEGIN;

CREATE TABLE IF NOT EXISTS public.certificates (
    id BIGSERIAL PRIMARY KEY,
    donation_id BIGINT REFERENCES public.blood_donations(id),
    user_id BIGINT REFERENCES public.users(id),
    certificate_number VARCHAR(255) UNIQUE,
    digital_signature TEXT,
    certificate_url VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

COMMIT;