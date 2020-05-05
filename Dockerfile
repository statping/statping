FROM alpine:latest
MAINTAINER "Hunter Long (https://github.com/hunterlong)"
ARG VERSION
ARG ARCH

RUN apk --no-cache add curl jq

RUN curl -L -s https://assets.statping.com/sass -o /usr/local/bin/sass && \
    chmod +x /usr/local/bin/sass

RUN curl -L -O https://github.com/statping/statping/releases/download/v$VERSION/statping-linux-$ARCH.tar.gz && \
    tar -xzf statping-linux-$ARCH.tar.gz && mv statping /usr/local/bin/ && rm -f statping-linux-$ARCH.tar.gz

WORKDIR /app
VOLUME /app

ENV PORT=8080
ENV STATPING_DIR=/app

EXPOSE $PORT

HEALTHCHECK --interval=30s --timeout=5s --retries=5 CMD curl -s "http://localhost:$PORT/health" | jq -r -e ".online==true"

CMD statping --port $PORT
