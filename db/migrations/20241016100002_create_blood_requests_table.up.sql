BEGIN;

CREATE TABLE IF NOT EXISTS public.blood_requests (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES public.users(id),
    hospital_id BIGINT REFERENCES public.hospitals(id),
    patient_name VARCHAR(255),
    blood_type VARCHAR(255),
    quantity INT,
    urgency_level VARCHAR(255),
    diagnosis TEXT,
    status VARCHAR(255),
    expiry_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

COMMIT;