FROM mysql:8.0

COPY migrations /migrations/
RUN cd migrations
RUN chmod +x migrate.sh

# Run the migrations and then start MySQL
CMD ["/bin/bash", "-c", "/migrate.sh && exec mysqld"]