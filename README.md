# SimplifiedCrafter

## Prerequisites

Before you begin, make sure you have the following installed:

- Docker
- Docker Compose

## Setup Instructions

### 1. Environment Variables

Create a `.env` file in the root directory of the project (next to `docker-compose.yml`) with the following content:

```ini
MYSQL_ROOT_PASSWORD=rootpassword
MYSQL_DATABASE=mydatabase
MYSQL_USER=myuser
MYSQL_PASSWORD=mypassword
MYSQL_HOST=localhost
MYSQL_PORT=3306
```

### 2. Running the Docker Compose Setup

Build and start the containers: Once your .env file is set up, you can start the MySQL service along with migrations using Docker Compose.

```bash
docker-compose up -d --build
```

This command:

* Builds the Docker images if they haven't been built yet.
* Starts the MySQL container and applies the migrations automatically.
* The migrations are executed in the order of the folder names (e.g., 001-initial_migration, 002-create_test-realm).