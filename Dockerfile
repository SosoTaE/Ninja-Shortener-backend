# Builder stage
FROM alpine:latest AS builder

# Install Go (replace 1.22 with the version in your go.mod)
RUN apk add --no-cache go

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod go.sum ./

# Download dependencies using Go Modules
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go binary (replace "url_shortener" with your project's main package name)
RUN go build -o /url_shortener

# Final stage
FROM alpine:latest

# Install dependencies for mkcert (including openssl and ca-certificates)
RUN apk add --no-cache bash openssl ca-certificates

# Install mkcert
RUN wget -qO /tmp/mkcert https://github.com/FiloSottile/mkcert/releases/download/v1.4.4/mkcert-v1.4.4-linux-amd64 && \
    chmod +x /tmp/mkcert && \
    mv /tmp/mkcert /usr/local/bin/mkcert

# Generate the certificate and key
RUN mkcert -install && \
    mkcert localhost 127.0.0.1 ::3000

# Set the working directory
WORKDIR /app

# Copy only the necessary binary from the builder stage
COPY --from=builder /url_shortener .

# Expose the port your application listens on (if applicable)
EXPOSE 443  # Change to 443 for HTTPS

# Command to run the application (using HTTPS)
CMD ["./url_shortener"]
