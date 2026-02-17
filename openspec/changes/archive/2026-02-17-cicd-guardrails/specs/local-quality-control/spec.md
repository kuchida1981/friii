## ADDED Requirements

### Requirement: `pre-commit` によるコミット前自動チェック
システムは `pre-commit` フレームワークを導入し、開発者がローカル環境で `git commit` を実行する際に、自動的に以下のチェックを実行しなければならない。
- 静的解析 (Lint) およびテスト
- OpenSpec の整合性検証 (外部提供)
- イシュー紐付けの検証 (外部提供)

#### Scenario: 外部提供チェックの失敗によるブロック
- **WHEN** OpenSpec のアーカイブが未完了の状態で `git commit` を実行する
- **THEN** 外部提供された検証スクリプトがエラーを返し、コミットが拒否される

### Requirement: 開発環境への `pre-commit` セットアップの自動化
システムは、開発者がプロジェクト参加時に容易にローカルガードレールを有効化できるよう、セットアップコマンドを提供しなければならない。

#### Scenario: `pre-commit` の初期化
- **WHEN** 開発者が `pre-commit install` を実行する
- **THEN** 以降の `git commit` 操作にフックが登録され、自動チェックが有効になる
