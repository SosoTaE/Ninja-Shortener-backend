# Use an official Go base image, selecting for Alpine for its smaller size.
FROM golang:1.20-alpine AS builder

# Set the working directory within the container.
WORKDIR /app

# Copy go.mod and go.sum to ensure dependencies are cached correctly.
COPY go.mod go.sum ./

# Download dependencies.
RUN go mod download

# Copy the entire project into the container.
COPY . .

# Build the Go binary statically for better portability.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Use a minimal base image for the final image. This minimizes the size of the container.
FROM alpine:latest

# Set the working directory for the run time container
WORKDIR /app

# Copy the binary from the builder stage to the final image.
COPY --from=builder /app/main .

# Define the command to run when the container starts.
CMD ["./main"]
