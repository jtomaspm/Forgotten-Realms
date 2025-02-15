#!/bin/bash

MYSQL_USER="${MYSQL_USER}"
MYSQL_PASSWORD="${MYSQL_PASSWORD}"
MYSQL_HOST="${MYSQL_HOST:-localhost}"
MYSQL_PORT="${MYSQL_PORT:-3306}"
MYSQL_DATABASE="${MYSQL_DATABASE}"

if [ -z "$MYSQL_USER" ] || [ -z "$MYSQL_PASSWORD" ] || [ -z "$MYSQL_HOST" ] || [ -z "$MYSQL_PORT" ] || [ -z "$MYSQL_DATABASE" ]; then
  echo "Error: Missing one or more MySQL environment variables."
  exit 1
fi

echo "Waiting for MySQL to start..."
for i in {1..60}; do  # Increase retries to 60
  if mysqladmin -h "$MYSQL_HOST" -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" -P "$MYSQL_PORT" ping --silent; then
    echo "MySQL is up!"
    break
  else
    echo "MySQL not yet ready... retrying ($i/60)"
    sleep 5 
  fi
done

if [ $i -eq 60 ]; then
  echo "MySQL didn't start in time, exiting."
  exit 1
fi

for folder in $(ls /migrations | sort); do
  echo "Running migrations from folder: $folder"

  for migration in $(ls /migrations/$folder/*.sql | sort); do
    echo "Running migration: $migration"

    mysql -h "$MYSQL_HOST" -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" -P "$MYSQL_PORT" "$MYSQL_DATABASE" < "/migrations/$folder/$(basename $migration)"

    if [ $? -eq 0 ]; then
      echo "Migration ($folder)[$migration] applied successfully."
    else
      echo "Error applying migration ($folder)[$migration]. Exiting."
      exit 1
    fi
  done
done