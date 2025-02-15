FROM mysql:9.2

EXPOSE ${MYSQL_PORT}

RUN microdnf install -y gettext

COPY migrations /migrations/
WORKDIR /migrations
RUN chmod +x migrate.sh

ENTRYPOINT ["bash", "-c", "docker-entrypoint.sh mysqld & ./migrate.sh && wait"]