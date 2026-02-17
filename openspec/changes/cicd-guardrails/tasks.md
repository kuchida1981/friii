## 1. Backend 品質ガードの構築

- [ ] 1.1 `.golangci.yml` を作成し、厳格な静的解析ルールを設定
- [ ] 1.2 `Makefile` またはシェルスクリプトを作成し、カバレッジ計測付きのテスト実行コマンドを用意
- [ ] 1.3 バックエンドの既存コードに対して Lint を実行し、修正が必要な箇所を対応

## 2. Frontend 品質ガードの構築

- [ ] 2.1 `.eslintrc.cjs` (または `eslint.config.js`) を更新し、React/Hooks 用の厳格なルールを追加
- [ ] 2.2 Vitest に `coverage-v8` を導入し、カバレッジ閾値を設定
- [ ] 2.3 フロントエンドの既存コードに対して Lint を実行し、修正が必要な箇所を対応

## 3. ローカルガード (pre-commit) の導入

- [ ] 3.1 プロジェクトルートに `.pre-commit-config.yaml` を作成
- [ ] 3.2 Backend/Frontend の Lint およびテストを hook として登録
- [ ] 3.3 ユーザー提供の OpenSpec/Issue チェック用スクリプトの呼び出し設定を追加
- [ ] 3.4 ローカルで `pre-commit install` を行い、正常に動作することを検証

## 4. CI パイプライン (GitHub Actions) の構築

- [ ] 4.1 `.github/workflows/ci.yml` を作成
- [ ] 4.2 プルリクエスト時に Backend/Frontend の Lint とテスト（カバレッジ計測含む）を実行するジョブを定義
- [ ] 4.3 ユーザー提供のチェック項目（OpenSpec, Issue, Todoリスト）を実行するジョブの枠組みを用意
- [ ] 4.4 カバレッジが閾値を下回った場合にジョブを失敗させるステップを追加

## 5. ドキュメントの更新

- [ ] 5.1 `DEVELOPMENT.md` に `pre-commit` の導入手順と CI ガードレールの説明を追記
