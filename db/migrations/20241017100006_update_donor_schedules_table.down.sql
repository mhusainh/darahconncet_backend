BEGIN;

-- Drop columns that were added
ALTER TABLE public.donor_schedules
DROP COLUMN IF EXISTS user_id,
DROP COLUMN IF EXISTS request_id;

COMMIT;