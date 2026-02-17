-- Seed Account Items using Category lookup
DO $$
DECLARE
    cat_assets_cur UUID := (SELECT id FROM account_categories WHERE name = '流動資産');
    cat_assets_fix UUID := (SELECT id FROM account_categories WHERE name = '固定資産');
    cat_liab_cur   UUID := (SELECT id FROM account_categories WHERE name = '流動負債');
    cat_equity     UUID := (SELECT id FROM account_categories WHERE name = '純資産');
    cat_rev_main   UUID := (SELECT id FROM account_categories WHERE name = '売上高');
    cat_rev_ext    UUID := (SELECT id FROM account_categories WHERE name = '営業外収益');
    cat_exp_admin  UUID := (SELECT id FROM account_categories WHERE name = '販売費及び一般管理費');
BEGIN
    INSERT INTO account_items (id, category_id, code, name, valid_from, created_at) VALUES
        -- 資産
        (gen_random_uuid(), cat_assets_cur, '100', '普通預金', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_assets_cur, '110', '売掛金', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_assets_fix, '150', '工具器具備品', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_assets_fix, '160', '長期前払費用', '2026-01-01', CURRENT_TIMESTAMP),
        
        -- 負債
        (gen_random_uuid(), cat_liab_cur, '200', '未払金', '2026-01-01', CURRENT_TIMESTAMP),
        
        -- 純資産 (事業主勘定)
        (gen_random_uuid(), cat_equity, '300', '元入金', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_equity, '310', '事業主借', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_equity, '320', '事業主貸', '2026-01-01', CURRENT_TIMESTAMP),
        
        -- 収益
        (gen_random_uuid(), cat_rev_main, '400', '売上', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_rev_ext, '450', '雑収入', '2026-01-01', CURRENT_TIMESTAMP),
        
        -- 費用
        (gen_random_uuid(), cat_exp_admin, '500', '旅費交通費', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '510', '地代家賃', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '520', '水道光熱費', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '530', '通信費', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '540', '消耗品費', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '550', '接待交際費', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '560', '損害保険料', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '570', '減価償却費', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '580', '修繕費', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '590', '長期前払費用償却', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '591', '支払手数料', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '592', '新聞図書費', '2026-01-01', CURRENT_TIMESTAMP),
        (gen_random_uuid(), cat_exp_admin, '599', '雑費', '2026-01-01', CURRENT_TIMESTAMP);
END $$;
