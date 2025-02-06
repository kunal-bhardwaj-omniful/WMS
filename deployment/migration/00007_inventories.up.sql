CREATE TABLE inventories (
                             id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                             sku_id uuid NOT NULL,
                             hub_id uuid NOT NULL,
                             available_qty integer NOT NULL DEFAULT 0,
                             allocated_qty integer NOT NULL DEFAULT 0,
                             damaged_qty integer NOT NULL DEFAULT 0,
                             zone varchar(50),
                             rack varchar(50),
                             bin varchar(50),
                             min_threshold integer DEFAULT 0,
                             max_threshold integer DEFAULT 0,
                             last_counted_at timestamptz,
                             created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                             updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                             CONSTRAINT inventories_sku_hub_unique UNIQUE (sku_id, hub_id),
                             CONSTRAINT fk_inventories_sku FOREIGN KEY (sku_id)
                                 REFERENCES skus(id) ON DELETE RESTRICT,
                             CONSTRAINT fk_inventories_hub FOREIGN KEY (hub_id)
                                 REFERENCES hubs(id) ON DELETE RESTRICT,
                             CONSTRAINT check_qty_positive CHECK (available_qty >= 0 AND allocated_qty >= 0 AND damaged_qty >= 0)
);

CREATE INDEX idx_inventories_sku_id ON inventories(sku_id);
CREATE INDEX idx_inventories_hub_id ON inventories(hub_id);

CREATE TRIGGER update_inventories_updated_at
    BEFORE UPDATE ON inventories
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();