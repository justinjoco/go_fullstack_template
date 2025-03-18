#!/bin/sh
#!/bin/sh
echo "Waiting for postgres..."

while ! nc -z "postgres" "5432"; do
    sleep 0.1
done

echo "PostgreSQL started"

migrate -path /app/migrations -database "postgres://admin:password@postgres/app_db?sslmode=disable" up
exec /bin/app