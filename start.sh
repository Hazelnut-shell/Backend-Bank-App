#!/bin/sh

# if any command returns non-zero, the script will exit immediately
set -e

source /app/app.env

echo "run db migration"
# -path, path of migration files
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up


echo "start the app"
# take all parameters passed to the script and run it
exec "$@"
