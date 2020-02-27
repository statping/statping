FROM golang:1.13.5-alpine as base
LABEL maintainer="Hunter Long (https://github.com/hunterlong)"
ARG VERSION
RUN apk add --no-cache libstdc++ gcc g++ make git ca-certificates linux-headers wget curl jq
RUN curl -L -s https://assets.statping.com/sass -o /usr/local/bin/sass && \
    chmod +x /usr/local/bin/sass
WORKDIR /go/src/github.com/hunterlong/statping
ADD . /go/src/github.com/hunterlong/statping
RUN go mod vendor
RUN make dev-deps
RUN make install

ENV IS_DOCKER=true
ENV STATPING_DIR=/app
WORKDIR /app

CMD ["statping"]
