FROM mysql:latest
ENV AUTO_RUN_DIR /docker-entrypoint-initdb.d
ENV INSTALL_DB_SQL schema-prod.sql
COPY ./$INSTALL_DB_SQL $AUTO_RUN_DIR/
RUN chmod a+x $AUTO_RUN_DIR/$INSTALL_DB_SQL

