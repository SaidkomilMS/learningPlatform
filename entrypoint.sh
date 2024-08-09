#!/bin/sh
# Run migrations
goose -dir migrations postgres "user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME host=$DB_HOST sslmode=disable " up

# Start main application
exec "$@"
