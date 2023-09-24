# Use an official Golang runtime as a base image
FROM golang:1.21.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Generate GraphQL code
RUN go run github.com/99designs/gqlgen generate

# Build the Go application
RUN go build -o server .

# Use a lightweight Alpine Linux as the base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to this stage
COPY --from=builder /app/server .

# Expose the port your application is running on
EXPOSE 8080

# Start the Go application
CMD ["./server"]