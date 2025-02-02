# Build stage
FROM golang:1.23.5 AS builder

WORKDIR /app

# Copy Go module files first
COPY go.mod ./

# Download dependencies
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the Go binary for Linux ARM64 (IMPORTANT for Apple Silicon)
RUN GOOS=linux GOARCH=arm64 go build -o books-app

# Use a minimal Debian-based image for better compatibility
FROM debian:latest

WORKDIR /root/

# Copy the built Go binary from the builder stage
COPY --from=builder /app/books-app .

# Ensure the binary has execution permissions
RUN chmod +x books-app

# Expose port 8080
EXPOSE 8084

# Run the application
CMD ["./books-app"]