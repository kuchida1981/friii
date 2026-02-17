CREATE TABLE IF NOT EXISTS partners (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    registration_number TEXT, -- 適格請求書発行事業者番号 (T+13桁)
    valid_from TIMESTAMP WITH TIME ZONE NOT NULL,
    valid_to TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_partners_valid_period ON partners (valid_from, valid_to);
