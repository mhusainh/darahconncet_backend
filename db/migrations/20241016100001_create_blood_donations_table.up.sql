BEGIN;

CREATE TABLE IF NOT EXISTS public.blood_donations (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES public.users(id),
    hospital_id BIGINT REFERENCES public.hospitals(id),
    registration_id BIGINT,
    donation_date TIMESTAMPTZ,
    blood_type VARCHAR(255),
    status VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

COMMIT;