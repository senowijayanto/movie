FROM mysql:latest

ENV MYSQL_DATABASE movie_db
ENV MYSQL_ROOT_PASSWORD supersecret

COPY ./sql-scripts/ /docker-entrypoint-initdb.d/

EXPOSE 3306
