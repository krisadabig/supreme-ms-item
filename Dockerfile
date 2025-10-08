# Build stage
FROM golang:1.25.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd  # amd64 for x86; arm64 for ARM

# Runtime stage (scratch for minimal)
FROM scratch
COPY --from=builder /app/main /main
COPY .env .
COPY config.yaml .
EXPOSE 8080
CMD ["/main"]