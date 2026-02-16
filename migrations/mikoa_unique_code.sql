-- Enforce unique mkoa code (case-insensitive) for non-deleted rows.
-- Run from project root: psql -U postgres -d library -f migrations/mikoa_unique_code.sql
CREATE UNIQUE INDEX IF NOT EXISTS mikoa_code_unique_active
  ON public.mikoa (LOWER(code))
  WHERE deleted_at IS NULL;
