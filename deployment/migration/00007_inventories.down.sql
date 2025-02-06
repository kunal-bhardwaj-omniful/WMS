DROP TRIGGER IF EXISTS update_inventories_updated_at ON inventories;
DROP INDEX IF EXISTS idx_inventories_hub_id;
DROP INDEX IF EXISTS idx_inventories_sku_id;
DROP TABLE IF EXISTS inventories;