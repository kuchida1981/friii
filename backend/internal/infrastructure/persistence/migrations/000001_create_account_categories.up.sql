CREATE TABLE IF NOT EXISTS account_categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    normal_side TEXT NOT NULL CHECK (normal_side IN ('DEBIT', 'CREDIT')),
    report_type TEXT NOT NULL CHECK (report_type IN ('BS', 'PL')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE account_categories IS '勘定科目カテゴリ';
COMMENT ON COLUMN account_categories.id IS 'カテゴリID';
COMMENT ON COLUMN account_categories.name IS 'カテゴリ名';
COMMENT ON COLUMN account_categories.normal_side IS '貸借属性 (DEBIT: 借方, CREDIT: 貸方)';
COMMENT ON COLUMN account_categories.report_type IS 'レポート種別 (BS: 貸借対照表, PL: 損益計算書)';
COMMENT ON COLUMN account_categories.created_at IS '作成日時';
