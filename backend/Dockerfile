FROM golang:1.23-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o /bin/app .

FROM alpine:3.21.3

COPY --from=builder /bin/app /bin/app

EXPOSE 5000
ENV GIN_MODE=release
# Run the Go application
CMD ["/bin/app"]