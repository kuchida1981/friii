## Context

現状のシステムは永続化層を持たず、インメモリまたはモックでの動作に留まっている。優良電子帳簿保存法の要件を満たすためには、仕訳データの不変性（物理的な更新・削除の禁止）と、訂正・削除の履歴保持をデータベースレベルで担保する必要がある。また、勘定科目等のマスターデータも、過去の帳簿の見読性を維持するために世代管理が必要である。

## Goals / Non-Goals

**Goals:**
- PostgreSQL 16 を使用した不変データモデルの構築。
- 仕訳（Header）と仕訳明細（Line）の 1:N 構造による複式簿記データの表現。
- 勘定科目、取引先、カテゴリマスタの導入と、それらの世代管理（valid_from/to）。
- データベースレベルでの `UPDATE/DELETE` を防止するガードレールの実装。
- `golang-migrate` による再現可能なマイグレーション基盤の確立。

**Non-Goals:**
- 複雑な決算書出力ロジックの実装（本フェーズではデータモデルに注力）。
- ユーザー認証・認可の実装（別 Issue で対応）。
- フロントエンドの UI 実装。

## Decisions

### 1. 仕訳データの不変性担保
- **選択**: PostgreSQL の `TRIGGER` を使用して、`journal_entries` および `journal_lines` テーブルへの `UPDATE` および `DELETE` 操作を例外として拒絶する。
- **理由**: アプリケーションコード（Go）での制御だけでなく、DBレベルで強制することで、直接的な SQL 操作によるデータ改ざんを防止し、優良電子帳簿の要件を高いレベルで充足するため。

### 2. マスターデータの世代管理（有効期間方式）
- **選択**: `valid_from` (timestamp) と `valid_to` (nullable timestamp) を持つ世代管理方式を採用。
- **理由**: 単なる `is_deleted` フラグではなく、有効期間を持たせることで、「ある時点の取引」が「その時点で有効だったマスター名称」と正しく紐付いていることを、過去に遡って正確に再現できるため。

### 3. カテゴリマスタによる属性の正規化
- **選択**: 勘定科目に直接「借方/貸方」や「BS/PL」を持たせず、`account_categories` マスタを介してこれらの属性を定義する。
- **理由**: 「流動資産」などのカテゴリ単位で属性を管理することで、AI Agent やユーザーが新しい科目を追加した際の属性設定ミスを構造的に防ぐため。

### 4. テーブル設計（概要）
- **journal_entries**: `id`, `transaction_date`, `description`, `original_entry_id` (修正前仕訳へのポインタ)。
- **journal_lines**: `id`, `entry_id`, `side` (DEBIT/CREDIT), `account_item_id`, `amount`, `partner_id`。
- **account_items**: `id`, `category_id`, `code`, `name`, `valid_from`, `valid_to`。
- **account_categories**: `id`, `name`, `normal_side`, `report_type` (BS/PL)。

## Risks / Trade-offs

- **[Risk] 赤黒処理によるデータ量の増加** → [Mitigation] 会計データは通常、検索インデックスを適切に張れば数百万件程度までは PostgreSQL で十分高速に処理可能。
- **[Risk] トリガーによるメンテナンス性の低下** → [Mitigation] トリガーの役割を「物理削除・更新の禁止」という単純なガードレールに限定し、ビジネスロジックは Go 側で実装することで複雑化を避ける。
- **[Risk] 世代管理マスタの検索複雑化** → [Mitigation] `valid_from <= ? AND (valid_to IS NULL OR valid_to > ?)` という標準的なクエリパターンをリポジトリ層でカプセル化する。
