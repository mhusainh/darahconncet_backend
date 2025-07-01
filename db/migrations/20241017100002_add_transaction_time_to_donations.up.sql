BEGIN;

ALTER TABLE public.donations
ADD COLUMN IF NOT EXISTS transaction_time TIMESTAMPTZ;

-- Change amount column type from INT to BIGINT
ALTER TABLE public.donations
ALTER COLUMN amount TYPE BIGINT;

COMMIT;