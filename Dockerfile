FROM golang:latest

WORKDIR /src

COPY ./account ./account
COPY ./clc_services ./clc_services
COPY ./game_database ./game_database

WORKDIR /src/account

RUN go mod download
RUN go mod tidy

RUN go build -o /game-server main.go 

RUN sh -c "rm -rf ./src"
RUN go clean -cache

EXPOSE 5000

ENTRYPOINT [ "/game-server" ]