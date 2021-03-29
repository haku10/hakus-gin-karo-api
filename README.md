# hakusGinKaroApi
Rest API for GIn by Golang

# reference
https://github.com/gin-gonic/gin

## migration
* into container
`docker exec -it コンテナID bash`
`cd /go`

* goose status
`goose status`
* create table
`goose up`
* add migration file
`goose create [マイグレーション名] sql`


# DynamoDBの初期構築(ローカル環境)
`endpoint_url="http://localhost:9603"`

# テーブル作成
```
aws dynamodb create-table \
  --cli-input-json file://db/dynamoDB/test_table.json \
  --billing-mode PAY_PER_REQUEST \
  --endpoint-url ${endpoint_url} >> result.txt
```
# テストデータ投入
```
aws dynamodb put-item --table-name test_table \
  --item '{"id": {"S": "1"}}' \
  --endpoint-url ${endpoint_url}
```

# ツールインストール (DB Browser for SQLite)
`brew install --cask db-browser-for-sqlite`
