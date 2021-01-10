FROM golang:latest

# イメージのアップグレード
RUN apt-get update
RUN apt-get upgrade -y

# 初期化
COPY . /go/src/
WORKDIR /go/src/

# ライブラリのインストール
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql
RUN go get -u github.com/pressly/goose/cmd/goose
RUN go get -u github.com/kelseyhightower/envconfig

CMD ["go", "run", "main.go"]
