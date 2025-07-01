BEGIN;

ALTER TABLE public.blood_donations
DROP COLUMN IF EXISTS public_id,
DROP COLUMN IF EXISTS url_file;

COMMIT;