CREATE TABLE skus (
                      id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                      seller_id uuid NOT NULL,
                      name varchar(100) NOT NULL,
                      code varchar(50) NOT NULL,
                      description varchar(500),
                      category varchar(100),
                      subcategory varchar(100),
                      brand varchar(100),
                      model varchar(100),
                      uom varchar(20) NOT NULL,
                      weight numeric(10,3),
                      dimensions jsonb,
                      created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                      updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                      CONSTRAINT skus_seller_code_unique UNIQUE (seller_id, code),
                      CONSTRAINT fk_skus_seller FOREIGN KEY (seller_id)
                          REFERENCES sellers(id) ON DELETE RESTRICT
);

CREATE INDEX idx_skus_seller_id ON skus(seller_id);
CREATE INDEX idx_skus_category ON skus(category);
CREATE INDEX idx_skus_dimensions ON skus USING gin (dimensions);

CREATE TRIGGER update_skus_updated_at
    BEFORE UPDATE ON skus
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();