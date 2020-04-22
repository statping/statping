FROM alpine:latest
LABEL maintainer="Hunter Long (https://github.com/hunterlong)"
ARG VERSION
ARG ARCH

RUN apk --no-cache add curl jq libsass linux-headers ca-certificates

RUN curl -L -s https://assets.statping.com/sass -o /usr/local/bin/sass && \
    chmod +x /usr/local/bin/sass

RUN curl -fsSL https://github.com/statping/statping/releases/download/v$VERSION/statping-linux-$ARCH.tar.gz -o statping.tar.gz && \
    tar -C /usr/local/bin -xzf statping.tar.gz && rm statping.tar.gz

WORKDIR /app

ENV IS_DOCKER=true
ENV STATPING_DIR=/app
ENV PORT=8080

EXPOSE $PORT

HEALTHCHECK --interval=60s --timeout=10s --retries=3 CMD curl -s "http://localhost:$PORT/health" | jq -r -e ".online==true"

CMD statping --port $PORT
