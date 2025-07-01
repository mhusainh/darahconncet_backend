BEGIN;

CREATE TABLE IF NOT EXISTS public.donor_registrations (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES public.users(id),
    schedule_id BIGINT,
    status VARCHAR(255),
    notes TEXT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

COMMIT;