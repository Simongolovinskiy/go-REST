FROM golang:1.22-alpine

WORKDIR /app

# Установка необходимых зависимостей для goose
RUN apk add --no-cache git

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

ENV PATH=$PATH:/go/bin

COPY . .

RUN go build -o main ./cmd/main.go

CMD ["./main"]