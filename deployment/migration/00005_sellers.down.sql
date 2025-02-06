DROP TRIGGER IF EXISTS update_sellers_updated_at ON sellers;
DROP INDEX IF EXISTS idx_sellers_tenant_id;
DROP TABLE IF EXISTS sellers;