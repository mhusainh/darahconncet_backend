BEGIN;

-- Rename request_id column back to schedule_id
ALTER TABLE public.donor_registrations
RENAME COLUMN request_id TO schedule_id;

COMMIT;