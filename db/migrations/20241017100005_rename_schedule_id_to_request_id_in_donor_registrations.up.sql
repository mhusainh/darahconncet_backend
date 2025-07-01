BEGIN;

-- Rename schedule_id column to request_id
ALTER TABLE public.donor_registrations
RENAME COLUMN schedule_id TO request_id;

COMMIT;