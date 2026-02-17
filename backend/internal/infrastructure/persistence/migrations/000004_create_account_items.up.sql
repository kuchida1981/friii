CREATE TABLE IF NOT EXISTS account_items (
    id UUID PRIMARY KEY,
    category_id UUID NOT NULL REFERENCES account_categories(id),
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    valid_from TIMESTAMP WITH TIME ZONE NOT NULL,
    valid_to TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_account_items_valid_period ON account_items (valid_from, valid_to);
CREATE INDEX idx_account_items_code ON account_items (code);

COMMENT ON TABLE account_items IS '勘定科目マスタ（世代管理）';
COMMENT ON COLUMN account_items.id IS '勘定科目ID';
COMMENT ON COLUMN account_items.category_id IS 'カテゴリID';
COMMENT ON COLUMN account_items.code IS '勘定科目コード';
COMMENT ON COLUMN account_items.name IS '勘定科目名';
COMMENT ON COLUMN account_items.valid_from IS '有効開始日時';
COMMENT ON COLUMN account_items.valid_to IS '有効終了日時（NULLは現在有効）';
COMMENT ON COLUMN account_items.created_at IS '作成日時';
