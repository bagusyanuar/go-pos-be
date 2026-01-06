-- =========================
-- MATERIAL IMAGES
-- =========================

CREATE TABLE material_images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    material_id UUID,
    image_group_id VARCHAR(100) NOT NULL,
    type VARCHAR(20) NOT NULL, -- 'original' atau 'thumbnail'
    url TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT fk_material_images_material
        FOREIGN KEY (material_id)
        REFERENCES materials(id)
        ON DELETE SET NULL
);

-- Optional: index for reporting / join (PK already indexed, tapi eksplisit OK)
CREATE INDEX idx_material_images_material_id
ON material_images(material_id);

CREATE INDEX idx_material_images_image_group_id
ON material_images(image_group_id);
