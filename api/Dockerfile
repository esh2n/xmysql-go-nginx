FROM golang:1.16

## 作業ディレクトリ
WORKDIR /app

# モジュール管理のファイルをコピー
COPY ./go.mod .
COPY ./go.sum .

# 外部パッケージのダウンロード
RUN go mod download

EXPOSE 3000
