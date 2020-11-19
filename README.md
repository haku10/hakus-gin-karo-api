# hakusGinKaroApi
Rest API for GIn by Golang

# reference
https://github.com/gin-gonic/gin

## migration
* into container
`docker exec -it コンテナID bash`
`cd /go`
* create table
`goose up`
* add migration file
`goose create [マイグレーション名] sql`
