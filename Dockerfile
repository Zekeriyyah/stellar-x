# Dockerfile
# Multi-stage build for Go app

# Stage 1: Dependencies
FROM golang:1.24-alpine AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Stage 2: Builder
FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git bash
COPY --from=deps /go/pkg/mod /go/pkg/mod
COPY . .
COPY scripts/ scripts/
RUN go build -v -o stellar-x ./cmd/server/main.go

# Stage 3: Final runtime image
FROM alpine:3.18
WORKDIR /app
RUN apk --no-cache add ca-certificates bash
RUN adduser -D -s /bin/bash appuser

COPY --from=builder /app/stellar-x .
COPY --from=builder /app/scripts/wait-for-db-and-migrate.sh .
RUN chmod +x wait-for-db-and-migrate.sh
RUN chown -R appuser:appuser /app
USER appuser
EXPOSE 8080
CMD ["./wait-for-db-and-migrate.sh"]