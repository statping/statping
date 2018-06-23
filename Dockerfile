FROM golang:1.10.3-alpine

RUN apk update && apk add git g++

WORKDIR $GOPATH/src/github.com/hunterlong/statup/

COPY . $GOPATH/src/github.com/hunterlong/statup/
RUN go get github.com/GeertJohan/go.rice/rice
RUN go get -d -v
RUN rice embed-go
RUN go install
WORKDIR /app
VOLUME /app

EXPOSE 8080

CMD ["/go/bin/statup"]