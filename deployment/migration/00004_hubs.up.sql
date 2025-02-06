CREATE TABLE hubs (
                      id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                      tenant_id uuid NOT NULL,
                      name varchar(100) NOT NULL,
                      code varchar(20) NOT NULL,
                      address varchar(255) NOT NULL,
                      city varchar(100),
                      state varchar(100),
                      country varchar(100),
                      pincode varchar(20),
                      location varchar(30),
                      created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                      updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                      CONSTRAINT hubs_tenant_code_unique UNIQUE (tenant_id, code),
                      CONSTRAINT fk_hubs_tenant FOREIGN KEY (tenant_id)
                          REFERENCES tenants(id) ON DELETE RESTRICT
);

CREATE INDEX idx_hubs_tenant_id ON hubs(tenant_id);

CREATE TRIGGER update_hubs_updated_at
    BEFORE UPDATE ON hubs
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
