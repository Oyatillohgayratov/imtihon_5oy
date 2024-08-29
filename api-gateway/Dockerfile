FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM gcr.io/distroless/base

COPY --from=builder /app/main /app/
COPY ./tls /app/tls
COPY ./.env /app/.env
COPY ./cmd/docs /app/docs

WORKDIR /app

ENTRYPOINT ["/app/main"]
