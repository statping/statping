FROM golang:1.13.5-alpine as base
LABEL maintainer="Hunter Long (https://github.com/hunterlong)"
ARG VERSION
RUN apk add --no-cache libstdc++ gcc g++ make git ca-certificates linux-headers wget curl jq libsass
RUN curl -L -s https://assets.statping.com/sass -o /usr/local/bin/sass && \
    chmod +x /usr/local/bin/sass
WORKDIR /go/src/github.com/hunterlong/statping
ADD Makefile go.mod /go/src/github.com/hunterlong/statping/
RUN go mod vendor && \
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
