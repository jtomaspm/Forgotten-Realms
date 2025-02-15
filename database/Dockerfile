FROM mysql:8.0

EXPOSE ${MYSQL_PORT}

COPY migrations /migrations/
WORKDIR /migrations
RUN chmod +x migrate.sh

ENTRYPOINT ["bash", "-c", "docker-entrypoint.sh mysqld & ./migrate.sh && wait"]