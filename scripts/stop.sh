#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <environment>"
  exit 1
fi

ENVIRONMENT=$1

mapfile -t configs < <(find "./infrastructure/docker-compose/" -mindepth 1 -maxdepth 1 -type d -exec basename {} \;)

if  [[ "$ENVIRONMENT" == "archive" ]]; then
  echo "Archive is not a valid environment."
  exit 1
fi

if [[ ! " ${configs[@]} " =~ " ${ENVIRONMENT} " ]]; then
  echo "Invalid environment. Available environments are:"
  for env in "${configs[@]}"; do
    if [[ $env == "archive" ]]; then
      continue
    fi
    echo " - $env"
  done
  exit 1
fi

if [[ $2 == "clean" ]]; then
  echo "Deleting volumes and containers for environment: $ENVIRONMENT"
  docker-compose --project-directory . --env-file "./infrastructure/docker-compose/${ENVIRONMENT}/.env" -f "./infrastructure/docker-compose/${ENVIRONMENT}/docker-compose.yaml" down -v
  exit 0
fi
docker-compose --project-directory . --env-file "./infrastructure/docker-compose/${ENVIRONMENT}/.env" -f "./infrastructure/docker-compose/${ENVIRONMENT}/docker-compose.yaml" down
exit 0