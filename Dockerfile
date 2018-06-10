FROM golang:alpine as builder

RUN apk update && apk add git

COPY . $GOPATH/src/github.com/hunterlong/fusioner/
WORKDIR $GOPATH/src/github.com/hunterlong/fusioner/
RUN go get github.com/GeertJohan/go.rice/rice
RUN go get -d -v
RUN rice embed-go
RUN go install

EXPOSE 8080

ENTRYPOINT fusioner