DROP TRIGGER IF EXISTS update_skus_updated_at ON skus;
DROP INDEX IF EXISTS idx_skus_dimensions;
DROP INDEX IF EXISTS idx_skus_category;
DROP INDEX IF EXISTS idx_skus_seller_id;
DROP TABLE IF EXISTS skus;