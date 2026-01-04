-- =========================
-- MATERIAL INVENTORIES
-- =========================

CREATE TABLE material_inventories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    material_id UUID,
    quantity NUMERIC(18,6) NOT NULL DEFAULT 0,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT fk_material_inventories_material
        FOREIGN KEY (material_id)
        REFERENCES materials(id)
        ON DELETE SET NULL
);

-- Optional: index for reporting / join (PK already indexed, tapi eksplisit OK)
CREATE INDEX idx_material_inventories_material_id
ON material_inventories(material_id);
