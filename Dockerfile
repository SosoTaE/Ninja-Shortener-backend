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

# Set the working directory
WORKDIR /app

# Copy only the necessary binary from the builder stage
COPY --from=builder /url_shortener .

# Expose the port your application listens on (if applicable)
EXPOSE 3000

# Command to run the application
CMD ["./url_shortener"]

