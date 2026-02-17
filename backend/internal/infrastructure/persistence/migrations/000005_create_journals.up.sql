CREATE TABLE IF NOT EXISTS journal_entries (
    id UUID PRIMARY KEY,
    transaction_date DATE NOT NULL,
    description TEXT NOT NULL,
    original_entry_id UUID REFERENCES journal_entries(id), -- 修正前仕訳へのポインタ
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS journal_lines (
    id UUID PRIMARY KEY,
    entry_id UUID NOT NULL REFERENCES journal_entries(id),
    side TEXT NOT NULL CHECK (side IN ('DEBIT', 'CREDIT')),
    account_item_id UUID NOT NULL REFERENCES account_items(id),
    amount DECIMAL(15, 2) NOT NULL,
    partner_id UUID REFERENCES partners(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_journal_entries_date ON journal_entries (transaction_date);
CREATE INDEX idx_journal_lines_entry_id ON journal_lines (entry_id);
CREATE INDEX idx_journal_lines_account_item_id ON journal_lines (account_item_id);
