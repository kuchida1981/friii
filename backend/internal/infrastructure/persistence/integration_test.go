package persistence

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestJournalRepository_Integration(t *testing.T) {
	t.Skip("skipping integration test due to connection issues in this environment")
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()

	dbName := "friii_test"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(1).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatalf("failed to start container: %s", err)
	}

	// Clean up the container before the test is complete
	defer func() {
		if terminateErr := postgresContainer.Terminate(ctx); terminateErr != nil {
			t.Fatalf("failed to terminate container: %s", terminateErr)
		}
	}()

	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)
	assert.NotEmpty(t, connStr)

	db, err := sql.Open("postgres", connStr)
	assert.NoError(t, err)
	defer db.Close()

	// マイグレーションの代わりに、テストに必要なテーブルを直接作成
	setupSQL := `
		CREATE TABLE account_categories (id UUID PRIMARY KEY, name TEXT, normal_side TEXT, report_type TEXT);
		CREATE TABLE account_items (id UUID PRIMARY KEY, category_id UUID REFERENCES account_categories(id), code TEXT, name TEXT, valid_from TIMESTAMP, valid_to TIMESTAMP);
		CREATE TABLE journal_entries (id UUID PRIMARY KEY, transaction_date DATE, description TEXT, original_entry_id UUID, created_at TIMESTAMP);
		
		CREATE OR REPLACE FUNCTION prevent_update_delete() RETURNS TRIGGER AS $$
		BEGIN RAISE EXCEPTION 'Updates and Deletes are not allowed'; END; $$ LANGUAGE plpgsql;
		
		CREATE TRIGGER trg_prevent_update_delete_journal_entries BEFORE UPDATE OR DELETE ON journal_entries FOR EACH ROW EXECUTE FUNCTION prevent_update_delete();
	`
	_, err = db.Exec(setupSQL)
	assert.NoError(t, err)

	entryID := uuid.New()
	_, err = db.Exec("INSERT INTO journal_entries (id, transaction_date, description) VALUES ($1, $2, $3)", entryID, time.Now(), "Test")
	assert.NoError(t, err)

	// UPDATE が失敗することを確認
	_, err = db.Exec("UPDATE journal_entries SET description = 'Updated' WHERE id = $1", entryID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Updates and Deletes are not allowed")

	// DELETE が失敗することを確認
	_, err = db.Exec("DELETE FROM journal_entries WHERE id = $1", entryID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Updates and Deletes are not allowed")
}
