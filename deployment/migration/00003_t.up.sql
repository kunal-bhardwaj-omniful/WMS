CREATE TABLE tenants (
                         id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                         name varchar(100) NOT NULL,
                         email varchar(100) NOT NULL,
                         gstin varchar(15),
                         created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                         updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                         CONSTRAINT tenants_name_unique UNIQUE (name),
                         CONSTRAINT tenants_email_unique UNIQUE (email)
);

CREATE TRIGGER update_tenants_updated_at
    BEFORE UPDATE ON tenants
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();