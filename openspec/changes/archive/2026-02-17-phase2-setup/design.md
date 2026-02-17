## Context

プロジェクト「friii」の設計フェーズが完了し、実装フェーズ（Phase 2）に移行します。AI Agent による効率的かつ高品質なコーディングを実現するため、明確なアーキテクチャ定義と強力なテスト基盤を備えた初期環境を構築する必要があります。

## Goals / Non-Goals

**Goals:**
- Docker Compose による、バックエンド、フロントエンド、DBが連携した開発環境の構築。
- Go による Clean Architecture の骨組み（ディレクトリ構造）の定義。
- GraphQL (gqlgen) による型安全な API 基盤のセットアップ。
- AI Agent がモックやテストコードを容易に生成・維持できるテスト基盤（moq, sqlmock, Testcontainers）の導入。
- ホットリロードが機能する開発ワークフローの実現。

**Non-Goals:**
- 各ドメイン（仕訳、勘定科目など）の具体的な業務ロジックの実装（これらは個別の OpenSpec チェンジで実施する）。
- E2E テスト（Cypress/Playwright等）の導入（現時点では非優先）。
- 本番環境向けのデプロイ設定（CI/CD パイプラインの初期化は含むが、インフラ構築は対象外）。

## Decisions

### 1. ディレクトリ構造: Clean Architecture
- **Rationale**: レイヤーを厳密に分離することで、ビジネスロジック（Usecase）を外部依存（DB, API）から隔離し、100%モックによるユニットテストを可能にするため。
- **Structure**:
  - `backend/internal/domain`: エンティティとリポジトリインターフェース。
  - `backend/internal/usecase`: ビジネスロジック。
  - `backend/internal/infrastructure`: DB実装、外部API。
  - `backend/internal/interface`: GraphQL リゾルバ、HTTP ハンドラ。

### 2. モックツール: `moq`
- **Rationale**: インターフェースを元にシンプルな構造体を生成する。AI Agent にとっても「関数を代入するだけ」という直感的なテストコードが書けるため、メンテナンス性が高い。

### 3. データベーステスト戦略: `sqlmock` + `Testcontainers`
- **Rationale**: 
  - `sqlmock`: リポジトリ層のユニットテストにおいて、SQLの組み立てやエラーハンドリングを高速に検証するため。
  - `Testcontainers`: 本物の PostgreSQL 上で「不変性（UPDATE禁止）」などの DB 制約が正しく機能することを保証するため。

### 4. GraphQL 実装: `gqlgen`
- **Rationale**: スキーマ駆動開発を強制し、Go の構造体と GraphQL スキーマの型安全な紐付けを自動化するため。

### 5. フロントエンド: Vite + React + TypeScript
- **Rationale**: 標準的なモダンスタックを採用し、開発の高速化と型安全性を確保するため。API 通信には Apollo Client または URQL を検討。

## Risks / Trade-offs

- **[Risk] Docker 越しのホットリロードの遅延** → **Mitigation**: Go では `air`、Vite では標準の HMR を使用し、適切にボリュームマウントとポーリング設定を調整する。
- **[Risk] Clean Architecture によるボイラープレートの増加** → **Mitigation**: ディレクトリ構造とインターフェースの書き方をパターン化し、AI Agent に生成を任せることで開発負荷を軽減する。
- **[Risk] Testcontainers によるテスト実行時間の増大** → **Mitigation**: `sqlmock` によるユニットテストをメインとし、`Testcontainers` を用いた結合テストは重要な整合性チェックに絞って実施する。
