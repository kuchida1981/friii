DROP TRIGGER IF EXISTS trg_prevent_update_delete_journal_lines ON journal_lines;
DROP TRIGGER IF EXISTS trg_prevent_update_delete_journal_entries ON journal_entries;
DROP FUNCTION IF EXISTS prevent_update_delete();
