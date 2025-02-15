#!/bin/bash

MYSQL_USER="${MYSQL_USER}"
MYSQL_PASSWORD="${MYSQL_PASSWORD}"
MYSQL_HOST="${MYSQL_HOST}"
MYSQL_PORT="${MYSQL_PORT}"
MYSQL_DATABASE="${MYSQL_DATABASE}"

if [ -z "$MYSQL_USER" ] || [ -z "$MYSQL_PASSWORD" ] || [ -z "$MYSQL_HOST" ] || [ -z "$MYSQL_PORT" ] || [ -z "$MYSQL_DATABASE" ]; then
  echo "Error: Missing one or more MySQL environment variables."
  exit 1
fi

for folder in $(ls /migrations | sort); do
  echo "Running migrations from folder: $folder"

  for migration in $(ls /migrations/$folder/*.sql | sort); do
    echo "Running migration: $migration"

    mysql -h "$MYSQL_HOST" -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" -P "$MYSQL_PORT" "$MYSQL_DATABASE" < "$migration"

    if [ $? -eq 0 ]; then
      echo "Migration $migration applied successfully."
    else
      echo "Error applying migration $migration. Exiting."
      exit 1
    fi
  done
done
