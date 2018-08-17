FROM golang:1.10.3-alpine
LABEL maintainer "Hunter Long (https://github.com/hunterlong)"

# Statup 'test' image for running a full test using the production environment
ENV VERSION=$(VERSION)

RUN apk add --no-cache libstdc++ gcc g++ make git ca-certificates linux-headers wget curl

RUN curl -L -s https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -o $GOPATH/bin/dep && \
    chmod +x $GOPATH/bin/dep

WORKDIR $GOPATH/src/github.com/hunterlong/statup

COPY . .

RUN wget -q https://assets.statup.io/sass && \
    chmod +x sass && \
    mv sass /usr/local/bin/sass

RUN dep ensure && \
    go get github.com/GeertJohan/go.rice/rice

RUN cd cmd && go build -ldflags "-X main.VERSION=$VERSION" && mv cmd /usr/local/bin/statup

ENV IS_DOCKER=true

WORKDIR /app
ADD ./Makefile .
EXPOSE 8080
ENTRYPOINT statup