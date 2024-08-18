FROM golang:latest

WORKDIR /deps

COPY ./go.mod .
COPY ./main.go .

RUN go get github.com/CiaranOtter/command_line_clue/server
RUN go mod tidy

RUN go build -o /server main.go 

RUN sh -c "rm -rf ./src"
RUN go clean -cache -modcache -fuzzcache

EXPOSE 5000

ENTRYPOINT [ "/server" ]