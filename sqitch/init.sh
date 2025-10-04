#!/bin/sh
# wait-for-it.sh

set -e

cd /sqitch

host="$1"
port="$2"
echo $host
echo $port

until pg_isready -h "$host" -p "$port"; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - Initializing Database"

# Deploy Sqitch migrations
sqitch deploy db:pg://rthomas:rthomas@db:5432/qualitinvest_db

# Check if a backup file exists
if [ -f "/sqitch/backup.sql" ]; then
    echo "Restore the backup..."
    psql -h db -U youruser -d yourdb -f /sqitch/backup.sql
else
    echo "No backup found. Nothing to restore."
fi

# Display the status of migrations
sqitch status db:pg://rthomas:rthomas@db:5432/qualitinvest_db
