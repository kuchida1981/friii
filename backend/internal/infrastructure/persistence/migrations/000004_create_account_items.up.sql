CREATE TABLE IF NOT EXISTS account_items (
    id UUID PRIMARY KEY,
    category_id UUID NOT NULL REFERENCES account_categories(id),
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    valid_from TIMESTAMP WITH TIME ZONE NOT NULL,
    valid_to TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_account_items_valid_period ON account_items (valid_from, valid_to);
CREATE INDEX idx_account_items_code ON account_items (code);
