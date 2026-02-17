## Why

プロジェクトの全体設計が完了したため、実際の開発に向けた「Phase 2: 基盤構築」を開始します。本変更では、バックエンド（Go）とフロントエンド（React）が協調して動作し、かつ高いテスト容易性（Testability）を維持できる開発環境を構築します。

## What Changes

- **開発環境の構築**: Docker Compose を用いて、PostgreSQL、Go バックエンド、React フロントエンドが連携する環境を構築します。
- **バックエンド基盤**: Clean Architecture に基づいたディレクトリ構造を初期化し、GraphQL (gqlgen) をセットアップします。
- **フロントエンド基盤**: Vite + TypeScript + React を初期化し、バックエンドとの通信（GraphQL）の準備を行います。
- **テスト基盤の導入**: AI Agent が高品質なコードとテストを生成できるよう、`moq` (モック生成)、`sqlmock` (DBユニットテスト)、`Testcontainers` (DB結合テスト) を導入します。

## Capabilities

### New Capabilities
- `environment-setup`: Docker Compose によるローカル開発環境と、各スタック（Go/React）の初期プロジェクト構造の定義。
- `testing-infrastructure`: AI Agent による自動テスト生成を支える、モックおよび DB テストツールの構成と利用パターンの定義。

### Modified Capabilities
<!-- 既存の要件自体に変更はないため、空欄とします -->

## Impact

- `backend/`: Go プロジェクトの新規作成と構成。
- `frontend/`: React プロジェクトの新規作成と構成。
- `docker-compose.yml`: プロジェクトルートへの追加。
- `openspec/specs/`: 新しいケイパビリティ（setup, testing）の追加。
