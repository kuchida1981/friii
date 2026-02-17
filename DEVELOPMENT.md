# friii 開発者ガイド (Phase 2)

## 開発環境の起動
```bash
docker compose up -d
```
- **Backend**: http://localhost:8080 (GraphQL Playground: http://localhost:8080)
- **Frontend**: http://localhost:5173
- **Database**: localhost:5432 (User: user, Pass: password, DB: friii)

## データベース操作 (PostgreSQL)

### DB への直接接続 (psql)
```bash
docker compose exec db psql -U user -d friii
```

### マイグレーション (golang-migrate)
`docker compose up` 時に自動適用されます。手動操作は以下の通り：
- **適用 (Up)**: `docker compose run --rm migrate`
- **新規作成**: `docker compose run --rm backend migrate create -ext sql -dir internal/infrastructure/persistence/migrations -seq <name>`

## バックエンド開発 (Go)

### アーキテクチャ (Clean Architecture)
- `internal/domain`: ビジネスルール、エンティティ、リポジトリのインターフェース。
- `internal/usecase`: アプリケーション固有のビジネスロジック。
- `internal/infrastructure`: データベース実装 (PostgreSQL) や外部サービス。
- `internal/interface`: GraphQL リゾルバや HTTP ハンドラ。

### GraphQL スキーマの更新
1. `backend/api/schema.graphqls` を編集。
2. 以下のコマンドを実行してコードを再生成：
```bash
docker compose run backend go run github.com/99designs/gqlgen generate
```

### テストの実行
- **ユニットテスト (sqlmock等)**:
```bash
docker compose run backend go test ./...
```
- **結合テスト (Testcontainers)**:
※ ローカル環境で Docker-in-Docker が利用可能な場合に実行。

### モックの生成
インターフェースに `//go:generate moq ...` を記載し、以下を実行：
```bash
docker compose run backend go generate ./...
```

## フロントエンド開発 (React)

### テストの実行
```bash
cd frontend && npm test
```

## 品質管理ガイドライン

### ローカルガードレール (pre-commit)
コミット前に Lint やテストを自動実行するために `pre-commit` を導入しています。

1. **インストール**:
   ホストマシンに `pre-commit` をインストールします。
   ```bash
   pip install pre-commit  # または brew install pre-commit
   ```
2. **セットアップ**:
   プロジェクトルートでフックを有効化します。
   ```bash
   pre-commit install
   ```
3. **実行**:
   `git commit` 時に自動で走りますが、手動で全ファイルに実行することも可能です。
   ```bash
   pre-commit run --all-files
   ```

### テストカバレッジの強制
CI/CD およびローカルチェックにおいて、以下のカバレッジ閾値を設けています。
- **全体**: 90% 以上
- **コアロジック (internal/domain, internal/usecase)**: 100%

### CI/CD ガードレール (GitHub Actions)
プルリクエスト（Draft を除く）に対して以下の自動チェックが走ります。
- **Backend**: Lint (`golangci-lint`), Unit/Integration Tests, カバレッジ計測。
- **Frontend**: Lint (`ESLint`), Tests, カバレッジ計測。
- **Other**: OpenSpec バリデーション, Issue 紐付けチェック, Todo リスト未完了チェック。

これらがパスしない限り、マージすることはできません。

## AI Agent への指示ガイドライン
- 新しい機能を実装する際は、まず `domain` にインターフェースを定義し、`moq` でモックを生成してから `usecase` のテストを書くように指示してください。
- データベース操作は `sqlmock` を用いて正常系・異常系の両方をテストするようにしてください。
- カバレッジが 100% (Core) / 90% (Overall) を維持しているか確認するように指示してください。
