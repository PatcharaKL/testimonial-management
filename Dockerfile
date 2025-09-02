# Start from the official Golang image
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git (for go mod) and build tools
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the app
RUN go build -o testimonial-app app/main.go

# Final image
FROM alpine:latest
WORKDIR /app

# Install ca-certificates for HTTPS
RUN apk add --no-cache ca-certificates

# Copy built binary and migrations
COPY --from=builder /app/testimonial-app .
COPY migrations ./migrations

EXPOSE 3000

CMD ["./testimonial-app"]
