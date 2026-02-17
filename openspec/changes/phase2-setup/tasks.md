## 1. 開発環境の基盤構築

- [ ] 1.1 プロジェクトルートに `docker-compose.yml` を作成 (PostgreSQL, Backend, Frontend)
- [ ] 1.2 `backend/` ディレクトリを作成し、`go mod init` で初期化
- [ ] 1.3 `frontend/` ディレクトリを作成し、Vite + TypeScript + React を初期化

## 2. バックエンドのアーキテクチャ定義

- [ ] 2.1 `backend/internal/` 配下に `domain`, `usecase`, `infrastructure`, `interface` ディレクトリを作成
- [ ] 2.2 `backend/api/` に初期の GraphQL スキーマ (`schema.graphqls`) を作成
- [ ] 2.3 `gqlgen` を導入し、初期のリゾルバとコード生成をセットアップ
- [ ] 2.4 ホットリロード用の `air` 設定ファイル (`.air.toml`) を作成

## 3. テスト基盤のセットアップ

- [ ] 3.1 `moq` を導入し、`go generate` でモックを生成するワークフローを確立
- [ ] 3.2 `sqlmock` を導入し、リポジトリ層のユニットテストのサンプルを作成
- [ ] 3.3 `Testcontainers` を導入し、PostgreSQL 実機を用いた統合テストのサンプルを作成

## 4. フロントエンドの初期設定

- [ ] 4.1 GraphQL 通信ライブラリ (Apollo Client または urql) を導入
- [ ] 4.2 バックエンドとの接続確認用テスト画面を作成
- [ ] 4.3 `Vitest` および `React Testing Library` の初期設定

## 5. 動作確認とドキュメント更新

- [ ] 5.1 全サービスが Docker Compose 上で正常に起動し、連携できることを確認
- [ ] 5.2 開発用 README (実装者向けガイド) を作成
