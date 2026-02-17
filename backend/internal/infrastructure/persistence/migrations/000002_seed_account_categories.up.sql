-- Account Categories Initial Data
INSERT INTO account_categories (id, name, normal_side, report_type) VALUES
    (gen_random_uuid(), '流動資産', 'DEBIT', 'BS'),
    (gen_random_uuid(), '固定資産', 'DEBIT', 'BS'),
    (gen_random_uuid(), '流動負債', 'CREDIT', 'BS'),
    (gen_random_uuid(), '固定負債', 'CREDIT', 'BS'),
    (gen_random_uuid(), '純資産', 'CREDIT', 'BS'),
    (gen_random_uuid(), '売上高', 'CREDIT', 'PL'),
    (gen_random_uuid(), '売上原価', 'DEBIT', 'PL'),
    (gen_random_uuid(), '販売費及び一般管理費', 'DEBIT', 'PL'),
    (gen_random_uuid(), '営業外収益', 'CREDIT', 'PL'),
    (gen_random_uuid(), '営業外費用', 'DEBIT', 'PL');
