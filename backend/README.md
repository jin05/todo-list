# 環境構築手順

## Go 環境構築

```shell
# 依存ライブラリをインストール
$ go mod download
```

## DB 環境構築

1. MySQL5.7 をインストール。
2. 以下の make コマンドを順番に実行する。
   ```shell
   # schemalexをインストール
   $ make db_install_schemalex
   # migrationを実行
   $ make local db_migrate
   ```

## backend の起動

```shell
$ make local run
```

## ユニットテストの実行

```shell
$ make test
```

## ローカル確認

- backend と frontend を両方起動していると http://localhot:3000 で接続できるようになります。

## AWS デプロイ

- backend/build ブランチに push すると CodePipeline が発火するように設定しています。
