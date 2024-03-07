# Start from a minimal Docker image containing the Go runtime
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the application source code
COPY . .

# Build and run the Go application
CMD go build && go run main.go
