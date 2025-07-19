FROM golang:1.24.4 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR /app/cmd/fetcher
RUN CGO_ENABLED=0 GOOS=linux go build -o /fetcher

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /fetcher .
ENTRYPOINT ["./fetcher"]