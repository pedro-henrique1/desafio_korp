FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./ 
RUN go build -o http-server-projeto-korp .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/http-server-projeto-korp .

EXPOSE 8080

CMD ["./http-server-projeto-korp"]