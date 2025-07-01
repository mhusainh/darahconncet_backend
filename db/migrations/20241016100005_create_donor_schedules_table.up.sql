BEGIN;

CREATE TABLE IF NOT EXISTS public.donor_schedules (
    id BIGSERIAL PRIMARY KEY,
    hospital_id BIGINT REFERENCES public.hospitals(id),
    event_name VARCHAR(255),
    event_date TIMESTAMPTZ,
    start_time TIMESTAMPTZ,
    end_time TIMESTAMPTZ,
    slots_available INT,
    slots_booked INT,
    description TEXT,
    status VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

COMMIT;