DROP TRIGGER IF EXISTS update_hubs_updated_at ON hubs;
DROP INDEX IF EXISTS idx_hubs_tenant_id;
DROP TABLE IF EXISTS hubs;