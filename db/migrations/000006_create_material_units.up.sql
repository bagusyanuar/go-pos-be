-- Create material_units pivot table
CREATE TABLE material_units (
    material_id UUID NOT NULL,
    unit_id UUID NOT NULL,

    PRIMARY KEY (material_id, unit_id),

    -- Optional: tambahan field
    conversion_rate NUMERIC(10,4) NOT NULL , -- misal: 1 dus = 12 pcs
    is_default BOOLEAN DEFAULT FALSE,

    -- Foreign keys
    CONSTRAINT fk_material_units_material
        FOREIGN KEY (material_id)
        REFERENCES materials(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_material_units_unit
        FOREIGN KEY (unit_id)
        REFERENCES units(id)
        ON DELETE RESTRICT
);

-- Indexes for performance
CREATE INDEX idx_material_units_material_id ON material_units(material_id);
CREATE INDEX idx_material_units_unit_id ON material_units(unit_id);
