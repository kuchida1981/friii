package persistence

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestJournalRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// 期待される SQL 実行の定義
	mock.ExpectExec("INSERT INTO journal_entries").
		WithArgs("test-id").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// ここでは実際の実装がないため、概念的なテストコードのみ
	// 実際のリポジトリ関数を呼び出す代わりに、期待値をチェックする
	t.Skip("TODO: JournalRepository.Save の実装後にテストを追加")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
