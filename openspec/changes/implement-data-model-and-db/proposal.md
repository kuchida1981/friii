## Why

現状のシステムには永続化層が存在せず、会計データの核心である「仕訳」や「勘定科目」を保存・管理する仕組みがありません。
優良な電子帳簿保存法（75万円控除）の要件を満たすためには、訂正・削除履歴の厳格な保持とデータの整合性確保が不可欠であり、不変性（Immutability）を前提としたデータモデルの構築が急務です。

## What Changes

- **仕訳データの不変永続化**: 1:N（ヘッダー:明細）構造による仕訳保存機能の導入。物理削除・更新を禁止。
- **マスターデータの世代管理**: 勘定科目、取引先、カテゴリマスタの導入。有効期限による履歴管理。
- **不変性ガードレールの導入**: データベースレベルおよびドメイン層での `UPDATE/DELETE` 制限。
- **マイグレーション基盤の構築**: `golang-migrate` によるバージョン管理されたスキーマ変更。

## Capabilities

### New Capabilities
- `journal-entry-persistence`: 仕訳および仕訳明細の不変永続化と、赤黒処理（訂正仕訳）による履歴追跡。
- `master-data-versioning`: 勘定科目、取引先、カテゴリマスタの有効期間（valid_from/to）による世代管理。

### Modified Capabilities
- `accounting-core`: 不変性の要件をデータモデル設計に反映。
- `excellent-book-compliance`: 訂正・削除履歴の物理的な保持要件を DB 設計で具体化。

## Impact

- **Backend**: `internal/infrastructure/persistence` 配下のリポジトリ実装。
- **Database**: PostgreSQL 16 への新規テーブル（entries, lines, accounts, partners, categories）の追加。
- **CI**: マイグレーションの自動適用と、不変性を担保するためのテストの追加。
