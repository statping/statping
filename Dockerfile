FROM golang:1.10.3-alpine

RUN apk update && apk add git g++

RUN curl -o /usr/local/bin/statup https://github.com/hunterlong/statup/releases/download/v0.18/statup-linux-x64

WORKDIR /app
VOLUME /app

EXPOSE 8080

CMD ["/go/bin/statup"]