#!/bin/bash

./scripts/stop.sh "$@"

if [ $? -ne 0 ]; then
  echo "Error stopping the environment. Exiting."
  exit 1
fi

docker-compose --project-directory . --env-file "./infrastructure/docker-compose/$1/.env" -f "./infrastructure/docker-compose/$1/docker-compose.yaml" up -d --build