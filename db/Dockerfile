FROM mariadb:latest

ENV MARIADB_ROOT_PASSWORD=rootpassword
ENV MARIADB_DATABASE=simple_login
ENV MARIADB_USER=app_user
ENV MARIADB_PASSWORD=app_password

COPY tables.sql /docker-entrypoint-initdb.d/

EXPOSE 3306
