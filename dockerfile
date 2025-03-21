# Use a Golang image for building the app
FROM golang:1.23.5 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the source code
COPY . .

# Ensure the binary is built for Linux and is statically linked
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Use a minimal Alpine image for the final build
FROM alpine:latest

# Install necessary certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Set the working directory in the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Ensure the binary has execution permissions
RUN chmod +x ./main

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
