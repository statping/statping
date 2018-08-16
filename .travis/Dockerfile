FROM golang:1.10.3-alpine

RUN apk add --no-cache libstdc++ gcc g++ make git ca-certificates linux-headers

RUN wget -q https://assets.statup.io/sass && \
    chmod +x sass && \
    mv sass /usr/local/bin/sass
RUN sass -v

WORKDIR $GOPATH/src/github.com/hunterlong/statup

COPY . .
RUN make deps
RUN make compile
RUN make install

ENV VERSION=$(VERSION)
ENV IS_DOCKER=true
ENV SASS=/usr/local/bin/sass
ENV ONLY_DB=sqlite
ENV STATUP_DIR=/go/src/github.com/hunterlong/statup
ENV GO_ENV=test

EXPOSE 8080
ENTRYPOINT SASS=$(SASS) VERSION=$(VERSION) make test