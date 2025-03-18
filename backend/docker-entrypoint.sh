#!/bin/sh
#!/bin/sh
echo "Waiting for postgres..."

while ! nc -z "postgres" "5432"; do
    sleep 0.1
done

echo "PostgreSQL started"

migrate -path /app/migrations -database "${DATABASE_URL}?sslmode=disable" up
exec /bin/app