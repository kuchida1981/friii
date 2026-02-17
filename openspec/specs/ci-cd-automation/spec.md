## Purpose
この仕様は、GitHub Actions を利用した継続的インテグレーション（CI）パイプラインの要件を定義し、自動テスト、静的解析、およびカバレッジ閾値の強制による品質管理を自動化する。

## Requirements

### Requirement: GitHub Actions によるプルリクエスト時の自動検証
システムは GitHub Actions を使用して、プルリクエストの作成および更新時に、バックエンドとフロントエンドの両方に対してテストと Lint を自動的に実行しなければならない (SHALL)。

#### Scenario: PR 作成時の自動テスト実行
- **WHEN** 開発者が GitHub 上でプルリクエストを作成または更新する
- **THEN** GitHub Actions ワークフローが起動し、全ての単体テスト、結合テスト、および静的解析（Lint）が実行される

### Requirement: テストカバレッジの閾値チェック
GitHub Actions のワークフローは、テスト実行結果のカバレッジを計測し、定義された閾値（全体 90%、コアロジック 100%）を下回る場合にビルドを失敗させなければならない (SHALL)。

#### Scenario: カバレッジ不足によるビルド失敗
- **WHEN** テストカバレッジが 90% 未満のコード、またはコアロジックで 100% 未満のコードを含む PR が提出される
- **THEN** GitHub Actions のジョブが失敗し、マージがブロックされる原因として通知される

### Requirement: 静的解析（Lint）エラーの検出
GitHub Actions のワークフローは、`golangci-lint` (Backend) および `ESLint` (Frontend) を実行し、コードスタイルの違反やポテンシャルなバグを検出しなければならない (SHALL)。

#### Scenario: Lint エラーによるビルド失敗
- **WHEN** 静的解析ルールに違反するコードを含む PR が提出される
- **THEN** GitHub Actions のジョブが失敗し、違反箇所がレポートされる
