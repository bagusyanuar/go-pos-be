-- Enable UUID support
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create materials table
CREATE TABLE materials (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    material_category_id UUID,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    image TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ,

    -- Foreign key constraint
    CONSTRAINT fk_materials_material_categories FOREIGN KEY (material_category_id)
        REFERENCES material_categories(id)
        ON DELETE SET NULL
);

-- Index for soft deletes
CREATE INDEX idx_materials_deleted_at ON materials(deleted_at);