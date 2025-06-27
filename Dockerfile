FROM golang:1.23.2-alpine

# Air のインストール（開発用途なので本番不要なら削除してもOK）
RUN apk add --no-cache git curl && \
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

WORKDIR /app

COPY . .

RUN go mod tidy

# ✅ 本番用バイナリをビルド（これが重要）
RUN go build -o main .

EXPOSE 8080

# ✅ ビルドしたバイナリを起動
CMD ["./main"]
