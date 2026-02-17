CREATE TABLE IF NOT EXISTS partners (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    registration_number TEXT, -- 適格請求書発行事業者番号 (T+13桁)
    valid_from TIMESTAMP WITH TIME ZONE NOT NULL,
    valid_to TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_partners_valid_period ON partners (valid_from, valid_to);

COMMENT ON TABLE partners IS '取引先マスタ（世代管理）';
COMMENT ON COLUMN partners.id IS '取引先ID';
COMMENT ON COLUMN partners.name IS '取引先名';
COMMENT ON COLUMN partners.registration_number IS '適格請求書発行事業者番号';
COMMENT ON COLUMN partners.valid_from IS '有効開始日時';
COMMENT ON COLUMN partners.valid_to IS '有効終了日時（NULLは現在有効）';
COMMENT ON COLUMN partners.created_at IS '作成日時';
