# 環境構築手順

- Backend のローカル環境構築手順
  - https://github.com/jin05/todo-list/tree/main/backend/README.md
- Frontend のローカル環境構築手順
  - https://github.com/jin05/todo-list/tree/main/frontend/README.md

# API ドキュメント

- Swagger で記述して、 [Redocly/redoc](https://github.com/Redocly/redoc) を使って自動生成しています。
  - todo-list/backend/document/todo-list-api.json を元に、  
    todo-list/backend/document/redoc-static.html を生成しています。

# 使用している言語など

## backend

- Go  
  シンプルで処理速度も早く堅牢な言語なので。(あとは好みです)

- Web Framework

  - gorilla  
    機能ごとに Git リポジトリが切られており、必要な機能のみインストールできるため。

    - mux
    - context
    - schema

- その他

  - lestrrat-go/jwx  
    jwt の verify で使っています。スターの多い jwt-go を使用しようと思いましたが、メンテされていないので。
  - gorm  
    ORM。他のライブラリを使ったことがないのですが、個人的に一番扱いやすいので。

- データベース
  - RDB(MySQL5.7)  
    DynamoDB も考えましたが、アクセス数も少なく小規模な API なので。

## frontend

- next.js
