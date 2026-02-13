-- Run this if mikoa table exists with old schema (id, name, code, created_at, updated_at only)
-- Adds: status, created_by, updated_by, deleted_at, deleted_by

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'mikoa' AND column_name = 'status') THEN
        ALTER TABLE public.mikoa ADD COLUMN status VARCHAR(50) DEFAULT 'active' NOT NULL;
    END IF;
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'mikoa' AND column_name = 'created_by') THEN
        ALTER TABLE public.mikoa ADD COLUMN created_by BIGINT;
    END IF;
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'mikoa' AND column_name = 'updated_by') THEN
        ALTER TABLE public.mikoa ADD COLUMN updated_by BIGINT;
    END IF;
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'mikoa' AND column_name = 'deleted_at') THEN
        ALTER TABLE public.mikoa ADD COLUMN deleted_at TIMESTAMP;
    END IF;
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'mikoa' AND column_name = 'deleted_by') THEN
        ALTER TABLE public.mikoa ADD COLUMN deleted_by BIGINT;
    END IF;
END $$;
