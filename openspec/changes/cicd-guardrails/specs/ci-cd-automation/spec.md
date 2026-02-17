## ADDED Requirements

### Requirement: GitHub Actions によるプルリクエスト時の自動検証
システムは GitHub Actions を使用して、プルリクエストの作成および更新時に、バックエンドとフロントエンドの両方に対してテストと Lint を自動的に実行しなければならない。

#### Scenario: PR 作成時の自動テスト実行
- **WHEN** 開発者が GitHub 上でプルリクエストを作成または更新する
- **THEN** GitHub Actions ワークフローが起動し、全ての単体テスト、結合テスト、および静的解析（Lint）が実行される

### Requirement: テストカバレッジの閾値チェック
GitHub Actions のワークフローは、テスト実行結果のカバレッジを計測し、定義された閾値（暫定 80%）を下回る場合にビルドを失敗させなければならない。

#### Scenario: カバレッジ不足によるビルド失敗
- **WHEN** テストカバレッジが 80% 未満のコードを含む PR が提出される
- **THEN** GitHub Actions のジョブが失敗し、マージがブロックされる原因として通知される

### Requirement: 静的解析（Lint）エラーの検出
GitHub Actions のワークフローは、`golangci-lint` (Backend) および `ESLint` (Frontend) を実行し、コードスタイルの違反やポテンシャルなバグを検出しなければならない。

#### Scenario: Lint エラーによるビルド失敗
- **WHEN** 静的解析ルールに違反するコードを含む PR が提出される
- **THEN** GitHub Actions のジョブが失敗し、違反箇所がレポートされる
