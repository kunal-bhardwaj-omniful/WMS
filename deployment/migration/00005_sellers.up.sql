CREATE TABLE sellers (
                         id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                         tenant_id uuid NOT NULL,
                         name varchar(50) NOT NULL,
                         code varchar(20) NOT NULL,
                         contact_person varchar(100),
                         email varchar(100),
                         phone varchar(20),
                         created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                         updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                         CONSTRAINT sellers_tenant_code_unique UNIQUE (tenant_id, code),
                         CONSTRAINT fk_sellers_tenant FOREIGN KEY (tenant_id)
                             REFERENCES tenants(id) ON DELETE RESTRICT
);

CREATE INDEX idx_sellers_tenant_id ON sellers(tenant_id);

CREATE TRIGGER update_sellers_updated_at
    BEFORE UPDATE ON sellers
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
