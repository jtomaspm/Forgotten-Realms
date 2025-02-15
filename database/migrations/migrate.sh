#!/bin/bash

# Setup
MYSQL_USER="root"
MYSQL_PASSWORD="${MYSQL_ROOT_PASSWORD}"
MYSQL_HOST="${MYSQL_HOST:-localhost}"
MYSQL_PORT="${MYSQL_PORT:-3306}"
MYSQL_DATABASE="${MYSQL_DATABASE}"

if [ -z "$MYSQL_USER" ] || [ -z "$MYSQL_PASSWORD" ] || [ -z "$MYSQL_HOST" ] || [ -z "$MYSQL_PORT" ] || [ -z "$MYSQL_DATABASE" ]; then
  echo "Error: Missing one or more MySQL environment variables."
  exit 1
fi

export MYSQL_PWD="$MYSQL_PASSWORD"

# Wait on SQL Server
echo "Waiting for MySQL to start..."
for i in {1..60}; do
  if mysqladmin -h "$MYSQL_HOST" -u "$MYSQL_USER" -P "$MYSQL_PORT" ping --silent; then
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

# Check Migration framework
echo "Checking if 'Migrations' table exists..."
TABLE_EXISTS=$(mysql -h "$MYSQL_HOST" -u "$MYSQL_USER" -P "$MYSQL_PORT" "$MYSQL_DATABASE" -e "SHOW TABLES LIKE 'Migrations';" | grep 'Migrations' > /dev/null; echo $?)

if [ $TABLE_EXISTS -ne 0 ]; then
  echo "'Migrations' table not found. Creating it now..."
  mysql -h "$MYSQL_HOST" -u "$MYSQL_USER" -P "$MYSQL_PORT" "$MYSQL_DATABASE" < /migrations/migrate.sql
  echo "'Migrations' table created."
fi

for database in $(find /migrations -maxdepth 1 -mindepth 1 -type d | sort); do
  DB_NAME=$(basename $database)

  DB_EXISTS=$(mysql -h "$MYSQL_HOST" -u "$MYSQL_USER" -P "$MYSQL_PORT" -e "SHOW DATABASES LIKE '$DB_NAME';" | grep "$DB_NAME" > /dev/null; echo $?)
  
  if [ $DB_EXISTS -ne 0 ]; then
    echo "Database '$DB_NAME' not found. Creating it now..."
    mysql -h "$MYSQL_HOST" -u "$MYSQL_USER" -P "$MYSQL_PORT" -e "CREATE DATABASE \`$DB_NAME\`;"

    DB_EXISTS_AFTER_CREATION=$(mysql -h "$MYSQL_HOST" -u "$MYSQL_USER" -P "$MYSQL_PORT" -e "SHOW DATABASES LIKE '$DB_NAME';" | grep "$DB_NAME" > /dev/null; echo $?)
    if [ $DB_EXISTS_AFTER_CREATION -ne 0 ]; then
      echo "Error: Database '$DB_NAME' was not created successfully. Exiting."
      exit 1
    else
      echo "Created $DB_NAME sucessfully!"
    fi
  fi

  for folder in $(find "$database" -maxdepth 1 -mindepth 1 -type d | sort); do
    for migration in $(ls $folder/*.sql | sort); do
      echo "Running migration: $migration"
      
      INSERT_QUERY="INSERT INTO Migrations (\`Name\`, \`Database\`) VALUES ('$migration', '$DB_NAME');"

      # Attempt to insert migration entry, skip migration if insert fails (duplicate)
      mysql -h "$MYSQL_HOST" -u "$MYSQL_USER" -P "$MYSQL_PORT" "$MYSQL_DATABASE" -e "$INSERT_QUERY" 2>/dev/null

      if [ $? -ne 0 ]; then
        echo "Migration [$migration] already applied. Skipping."
        continue
      fi

      mysql -h "$MYSQL_HOST" -u "$MYSQL_USER"  -P "$MYSQL_PORT" "$DB_NAME" < "$migration"

      if [ $? -eq 0 ]; then
        echo "Migration [$migration] applied successfully."
      else
        echo "Error applying migration [$migration]. Exiting."
        exit 1
      fi
    done
  done
done

echo "Done applying migrations!"