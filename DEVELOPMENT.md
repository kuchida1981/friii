# friii 開発者ガイド (Phase 2)

## 開発環境の起動
```bash
docker compose up -d
```
- **Backend**: http://localhost:8080 (GraphQL Playground: http://localhost:8080)
- **Frontend**: http://localhost:5173

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

## AI Agent への指示ガイドライン
- 新しい機能を実装する際は、まず `domain` にインターフェースを定義し、`moq` でモックを生成してから `usecase` のテストを書くように指示してください。
- データベース操作は `sqlmock` を用いて正常系・異常系の両方をテストするようにしてください。
