# Use the official Golang image as the base image
FROM golang:1.21-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o greeting-app .

# Use a lightweight Alpine image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/greeting-app .

# Copy the static files
COPY ./static ./static

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./greeting-app"]