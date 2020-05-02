FROM statping/statping:base AS base

# Statping main Docker image that contains all required libraries
FROM alpine:latest
RUN apk --no-cache add libgcc libstdc++ ca-certificates curl jq && update-ca-certificates

COPY --from=base /go/bin/statping /usr/local/bin/
COPY --from=base /usr/local/bin/sass /usr/local/bin/
COPY --from=base /usr/local/share/ca-certificates /usr/local/share/

WORKDIR /app

ENV IS_DOCKER=true
ENV SASS=/usr/local/bin/sass
ENV STATPING_DIR=/app
ENV PORT=8080

EXPOSE $PORT

HEALTHCHECK --interval=60s --timeout=10s --retries=3 CMD curl -s "http://localhost:$PORT/health" | jq -r -e ".online==true"

CMD statping --port $PORT
