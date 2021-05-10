# 環境構築手順

- 依存ライブラリをインストール
  ```shell
  $ yarn
  ```
- build 実行

  ```shell
  $ yarn build
  ```

- frontend のローカルサーバを起動

  ```shell
  $ yarn dev
  ```
  ※ backend もローカル起動していると、http://localhot:3000 で接続できるようになります。

- AWS デプロイ
  ```shell
  $ yarn deploy
  ```
  ※ serverless を使用しているので、npm などで事前にインストールする。
