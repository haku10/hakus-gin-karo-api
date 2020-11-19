FROM golang:latest

# ライブラリのインストール
RUN go get -u github.com/gin-gonic/gin

# git導入のフォルダ作成
WORKDIR /rest_api

# 初期化
RUN go mod init rest_api

ADD ./rest_api /rest_api

CMD ["go", "run", "main.go"]
