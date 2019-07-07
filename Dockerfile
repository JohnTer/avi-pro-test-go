FROM golang:latest

RUN mkdir -p /home/app

WORKDIR /home/app

COPY ./app /home/app

RUN go get -d ./

RUN apt-get update && apt-get install -y sqlite

RUN mkdir database && sqlite3 ./database/numbers.db < ./sql/create_table.sql

RUN go build

CMD ["./app", "8080"]

EXPOSE 8080
