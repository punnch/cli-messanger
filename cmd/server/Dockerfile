# Multi-stage Dockerfile
FROM golang:1.26.2-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/server/exe /app/cmd/server

FROM alpine:3.23.4 AS runner
WORKDIR /app
COPY --from=builder /app/cmd/server/exe /app/exe
EXPOSE 9000
CMD ["/app/exe"]
