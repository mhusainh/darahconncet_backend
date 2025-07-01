BEGIN;

-- Add columns that exist in entity but not in migration
ALTER TABLE public.donor_schedules
ADD COLUMN IF NOT EXISTS user_id BIGINT REFERENCES public.users(id),
ADD COLUMN IF NOT EXISTS request_id BIGINT REFERENCES public.blood_requests(id);

COMMIT;