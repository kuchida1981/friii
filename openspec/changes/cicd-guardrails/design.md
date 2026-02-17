## Context

プロジェクト基盤（Go + Clean Architecture, React + Vite）が構築された直後のフェーズであり、今後の機能開発に備えて品質担保の自動化が必要です。現在は手動でのテスト実行と基本的な Lint のみが存在し、カバレッジの強制やコミット前の自動チェックはありません。

## Goals / Non-Goals

**Goals:**
- CI/CD パイプライン（GitHub Actions）で全テストと Lint を自動実行する。
- バックエンドとフロントエンドの両方でテストカバレッジの閾値（暫定 80%）を設定し、下回る場合は CI を失敗させる。
- `pre-commit` ツールを導入し、開発者がローカルで問題を早期発見できるようにする。
- 既存の Lint ルールをより厳格にし、コードスタイルの統一を徹底する。

**Non-Goals:**
- 本番環境への自動デプロイ設定（このフェーズでは品質ガードのみを対象とする）。
- E2E テストの導入（単体・結合テストと静的解析に集中する）。

## Decisions

- **Backend Lint**: `golangci-lint` を採用。`.golangci.yml` を作成し、`gocritic`, `revive`, `govet`, `staticcheck` などの強力なリンターを有効化する。
- **Frontend Lint**: 既存の ESLint 設定をベースに、`eslint-plugin-react`, `eslint-plugin-react-hooks` などを厳格化し、Prettier との競合を避ける設定を行う。
- **Coverage Tool**: 
  - Backend: Go 標準の `-cover` プロファイルを使用。
  - Frontend: Vitest の `coverage-v8` を使用。
- **Local Guard**: `pre-commit` フレームワークを採用。`.pre-commit-config.yaml` をプロジェクトルートに配置し、以下のチェックを実行する：
  - **私が実装**: `golangci-lint`, `eslint`, `vitest` (related tests)。
  - **ユーザー提供を統合**: `openspec` のアーカイブ・バリデーションチェック、イシュー紐付けチェック。
  - ※ ToDo リストのチェックのみ、性質上 GitHub Actions 側でのみ実行する。
- **CI Platform**: GitHub Actions を使用。`.github/workflows/ci.yml` を作成。ローカルの `pre-commit` と重複するチェックを行い、最終的なガードレールとする。

## Risks / Trade-offs

- **[Risk] テスト実行時間の増加** → [Mitigation] `pre-commit` では変更されたファイルに関連するテストのみを実行する、あるいは CI で並列実行を行う。
- **[Risk] カバレッジの閾値による開発の停滞** → [Mitigation] 暫定的に 80% とし、必要に応じて除外ディレクトリ（generated code 等）を適切に設定する。
- **[Risk] pre-commit 導入のオーバーヘッド** → [Mitigation] 開発ガイド (`DEVELOPMENT.md`) にセットアップ手順を明記し、初回の `pre-commit install` を容易にする。
