# build da aplicacao
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

# baixa as dependencias e compila
RUN go get github.com/prometheus/client_golang@v1.17.0 && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o http-server-projeto-korp ./cmd/server

# imagem final menor
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/http-server-projeto-korp .

EXPOSE 8080

CMD ["./http-server-projeto-korp"]
