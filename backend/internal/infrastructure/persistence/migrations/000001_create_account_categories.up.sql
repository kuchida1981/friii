CREATE TABLE IF NOT EXISTS account_categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    normal_side TEXT NOT NULL CHECK (normal_side IN ('DEBIT', 'CREDIT')),
    report_type TEXT NOT NULL CHECK (report_type IN ('BS', 'PL')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
