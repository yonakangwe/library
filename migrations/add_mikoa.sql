-- Mikoa (regions) table - full schema matching entity
-- For existing incomplete table: DROP TABLE IF EXISTS mikoa; then run this.
CREATE TABLE IF NOT EXISTS public.mikoa (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(20) NOT NULL,
    status VARCHAR(50) DEFAULT 'active' NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by BIGINT,
    updated_by BIGINT,
    deleted_by BIGINT
);

ALTER TABLE public.mikoa OWNER TO postgres;
