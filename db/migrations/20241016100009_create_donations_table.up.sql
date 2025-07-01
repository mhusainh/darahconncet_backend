BEGIN;

CREATE TABLE IF NOT EXISTS public.donations (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES public.users(id),
    amount INT,
    status VARCHAR(20),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

COMMIT;