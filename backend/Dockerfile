# Stage 1: Build the Go binary
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app .

# Stage 2: Create the runtime image
FROM alpine:3.21.3

# Install necessary libraries for running Go binary (if needed)
RUN apk --no-cache add ca-certificates

# Set the working directory for the final image
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/app .

# Expose the port the app runs on (replace with your actual app's port)
EXPOSE 5000

# Run the Go application
CMD ["./app"]