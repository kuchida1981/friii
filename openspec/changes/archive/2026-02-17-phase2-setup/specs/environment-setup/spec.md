## ADDED Requirements

### Requirement: Docker Compose による統合開発環境の提供
システムは Docker Compose を使用して、バックエンド、フロントエンド、およびデータベースが相互に通信可能な統合開発環境を提供しなければならない。

#### Scenario: 開発環境の一括起動
- **WHEN** 開発者が `docker-compose up` コマンドを実行する
- **THEN** PostgreSQL, Backend (Go), Frontend (React) の全てのサービスが正常に起動し、相互に接続される

### Requirement: バックエンド (Go) の Clean Architecture 構成
バックエンドは Clean Architecture の原則に従い、`internal/` 配下に `domain`, `usecase`, `infrastructure`, `interface` の各レイヤーを分離して保持しなければならない。

#### Scenario: レイヤー間の依存関係の検証
- **WHEN** バックエンドのコード構造を確認する
- **THEN** `domain` レイヤーが他のどのレイヤーにも依存していないこと、および `usecase` が `domain` にのみ依存していることが確認できる

### Requirement: フロントエンド (React) の初期プロジェクト構成
フロントエンドは Vite と TypeScript を使用して構築され、バックエンドの GraphQL エンドポイントと通信するための基本設定を保持しなければならない。

#### Scenario: フロントエンドの初期表示
- **WHEN** フロントエンドサービスを起動しブラウザでアクセスする
- **THEN** Vite + React の初期画面が表示され、開発サーバーが正常に動作していることが確認できる
