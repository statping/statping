ARG BUILDPLATFORM
ARG GIT_COMMIT_HASH
FROM razorpay/statping:base_$GIT_COMMIT_HASH AS base
# Statping main Docker image that contains all required libraries
FROM alpine:latest
RUN apk --no-cache add libgcc libstdc++ ca-certificates curl jq && update-ca-certificates
RUN apk add --update tzdata

COPY --from=base /go/bin/statping /usr/local/bin/
COPY --from=base /root/sassc/bin/sassc /usr/local/bin/
COPY --from=base /usr/local/share/ca-certificates /usr/local/share/

WORKDIR /app
VOLUME /app

COPY --from=base /go/src/github.com/statping/statping/react/ ./react/

COPY --from=base /go/src/github.com/statping/statping/configs/*.yml ./configs/

ENV IS_DOCKER=true
ENV SASS=/usr/local/bin/sassc
ENV STATPING_DIR=/app
ENV PORT=80
ENV PROMETHEUS_PORT=9000
ENV TZ="Asia/Kolkata"

EXPOSE $PORT
EXPOSE $PROMETHEUS_PORT

HEALTHCHECK --interval=60s --timeout=10s --retries=3 CMD curl -s "http://localhost:$PORT/health" | jq -r -e ".online==true"

CMD statping --port $PORT
