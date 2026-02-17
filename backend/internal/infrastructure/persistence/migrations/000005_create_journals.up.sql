CREATE TABLE IF NOT EXISTS journal_entries (
    id UUID PRIMARY KEY,
    transaction_date DATE NOT NULL,
    description TEXT NOT NULL,
    original_entry_id UUID REFERENCES journal_entries(id), -- 修正前仕訳へのポインタ
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS journal_lines (
    id UUID PRIMARY KEY,
    entry_id UUID NOT NULL REFERENCES journal_entries(id),
    side TEXT NOT NULL CHECK (side IN ('DEBIT', 'CREDIT')),
    account_item_id UUID NOT NULL REFERENCES account_items(id),
    amount DECIMAL(15, 2) NOT NULL,
    partner_id UUID REFERENCES partners(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_journal_entries_date ON journal_entries (transaction_date);
CREATE INDEX idx_journal_lines_entry_id ON journal_lines (entry_id);
CREATE INDEX idx_journal_lines_account_item_id ON journal_lines (account_item_id);

COMMENT ON TABLE journal_entries IS '仕訳ヘッダー（不変）';
COMMENT ON COLUMN journal_entries.id IS '仕訳ID';
COMMENT ON COLUMN journal_entries.transaction_date IS '取引日';
COMMENT ON COLUMN journal_entries.description IS '取引説明';
COMMENT ON COLUMN journal_entries.original_entry_id IS '修正元仕訳ID（赤黒処理用）';
COMMENT ON COLUMN journal_entries.created_at IS '作成日時';

COMMENT ON TABLE journal_lines IS '仕訳明細（不変）';
COMMENT ON COLUMN journal_lines.id IS '明細ID';
COMMENT ON COLUMN journal_lines.entry_id IS '仕訳ID';
COMMENT ON COLUMN journal_lines.side IS '貸借属性 (DEBIT: 借方, CREDIT: 貸方)';
COMMENT ON COLUMN journal_lines.account_item_id IS '勘定科目ID';
COMMENT ON COLUMN journal_lines.amount IS '金額';
COMMENT ON COLUMN journal_lines.partner_id IS '取引先ID';
COMMENT ON COLUMN journal_lines.created_at IS '作成日時';
