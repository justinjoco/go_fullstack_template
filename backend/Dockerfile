FROM golang:1.23-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Install necessary dependencies for the PostgreSQL driver
RUN go install -tags 'postgres'  github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.2

# Build the Go application
RUN go build -o /bin/app .

FROM alpine:3.21.3 as release

WORKDIR /app

COPY --from=builder /app/docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /bin/app /bin/app

COPY migrations /app/migrations

EXPOSE 5000
ENV GIN_MODE=release
# Run the Go application

RUN chmod +x /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]