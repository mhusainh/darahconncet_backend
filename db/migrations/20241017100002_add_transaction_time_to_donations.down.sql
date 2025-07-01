BEGIN;

-- Drop the transaction_time column
ALTER TABLE public.donations
DROP COLUMN IF EXISTS transaction_time;

-- Change amount column type back to INT
ALTER TABLE public.donations
ALTER COLUMN amount TYPE INT;

COMMIT;