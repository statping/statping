FROM hunterlong/statping:base as base

FROM alpine:latest
LABEL maintainer="Hunter Long (https://github.com/hunterlong)"

ARG VERSION

RUN apk --no-cache add curl jq libsass

COPY --from=base /usr/local/bin/sass /usr/local/bin/sass
COPY --from=base /go/bin/statping /usr/local/bin/statping

WORKDIR /app
VOLUME /app

ENV IS_DOCKER=true
ENV STATPING_DIR=/app
ENV PORT=8080

EXPOSE $PORT

HEALTHCHECK --interval=60s --timeout=10s --retries=3 CMD curl -s "http://localhost:$PORT/health" | jq -r -e ".online==true"

CMD statping -port $PORT
