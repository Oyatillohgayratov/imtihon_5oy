FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd
RUN go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM gcr.io/distroless/base

COPY --from=builder /app/main /app/
COPY ./.env /app/.env

WORKDIR /app

ENTRYPOINT ["/app/main"]