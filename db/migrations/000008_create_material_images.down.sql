-- =========================
-- DROP MATERIAL IMAGES
-- =========================

DROP INDEX IF EXISTS idx_material_images_material_id;
DROP INDEX IF EXISTS idx_material_images_image_group_id;

DROP TABLE IF EXISTS material_images;
