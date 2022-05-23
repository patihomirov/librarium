FROM ubuntu:latest
MAINTAINER PATihomirov
LABEL version="1.0"
RUN apt-get update
RUN apt-get upgrade -y
#RUN apt-get install mysql-server -y
RUN apt-get install golang -y
RUN mkdir "/app"
ADD . /app
WORKDIR /app
#команды к mySQL не смог реализовать, т.к. docker устанавливает на базу данных случайный пароль и его потом не определить
#базу данных нужно развернуть отдельно и настроить связь с ней, либо использовать SQLite
#RUN mysql -u root -p passw0rd librarium_base < ./database_dump/librarium_base.sql
#RUN mysql librarium_base < ./database_dump/librarium_base.sql
#RUN mysql -u root -p passw0rd -e "GRANT ALL PRIVILEGES ON librarium_base TO 'web'@'localhost' IDENTIFIED BY '123' WITH GRANT OPTION;"
RUN go build -o main ./cmd
CMD ["/app/main"]