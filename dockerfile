FROM golang:1.16.0-alpine3.13 as build
WORKDIR /app

#FROM base as dev

RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

#RUN apk add --update --no-cache ca-certificates git

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

#RUN go get github.com/githubnemo/CompileDaemon

COPY . .

#COPY ./entrypoint.sh /entrypoint.sh

RUN go build -o main .

#ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
#RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

#ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o main main.go" -command="./main"

#ENTRYPOINT [ "sh", "/entrypoint.sh" ]

EXPOSE 3000

CMD ["./main"]