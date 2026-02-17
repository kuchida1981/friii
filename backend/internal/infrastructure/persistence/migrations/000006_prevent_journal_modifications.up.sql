-- Function to prevent UPDATE or DELETE
CREATE OR REPLACE FUNCTION prevent_update_delete()
RETURNS TRIGGER AS $$
BEGIN
    RAISE EXCEPTION 'Updates and Deletes are not allowed on this table for immutability.';
END;
$$ LANGUAGE plpgsql;

-- Apply to journal_entries
CREATE TRIGGER trg_prevent_update_delete_journal_entries
BEFORE UPDATE OR DELETE ON journal_entries
FOR EACH ROW EXECUTE FUNCTION prevent_update_delete();

-- Apply to journal_lines
CREATE TRIGGER trg_prevent_update_delete_journal_lines
BEFORE UPDATE OR DELETE ON journal_lines
FOR EACH ROW EXECUTE FUNCTION prevent_update_delete();
