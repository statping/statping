FROM golang:1.12-alpine as base
LABEL maintainer="Hunter Long (https://github.com/hunterlong)"
ARG VERSION
ENV DEP_VERSION v0.5.0
RUN apk add --no-cache libstdc++ gcc g++ make git ca-certificates linux-headers wget curl jq libsass
RUN curl -L -s https://github.com/golang/dep/releases/download/$DEP_VERSION/dep-linux-amd64 -o /go/bin/dep && \
    chmod +x /go/bin/dep
RUN curl -L -s https://assets.statping.com/sass -o /usr/local/bin/sass && \
    chmod +x /usr/local/bin/sass
WORKDIR /go/src/github.com/hunterlong/statping
ADD Makefile Gopkg.* /go/src/github.com/hunterlong/statping/
RUN make dep && \
    make dev-deps
ADD . /go/src/github.com/hunterlong/statping
RUN make install

# Statping :latest Docker Image
FROM alpine:latest
LABEL maintainer="Hunter Long (https://github.com/hunterlong)"

ARG VERSION
ENV IS_DOCKER=true
ENV STATPING_DIR=/app
ENV PORT=8080
RUN apk --no-cache add curl jq libsass

COPY --from=base /usr/local/bin/sass /usr/local/bin/sass
COPY --from=base /go/bin/statping /usr/local/bin/statping

WORKDIR /app
VOLUME /app
EXPOSE $PORT

HEALTHCHECK --interval=5s --timeout=5s --retries=5 CMD curl -s "http://localhost:$PORT/health" | jq -r -e ".online==true"

CMD statping -port $PORT
