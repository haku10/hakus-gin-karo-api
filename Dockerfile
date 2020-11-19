FROM golang:latest

# ライブラリのインストール
RUN go get -u github.com/gin-gonic/gin

# git導入のフォルダ作成
WORKDIR /go/src/rest_api

# 初期化
RUN go mod init rest_api

ADD ./rest_api /go/src/rest_api
ADD ./db go/db

CMD ["go", "run", "main.go"]
