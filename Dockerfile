ARG VERSION
FROM hunterlong/statup:base-v${VERSION}
MAINTAINER "Hunter Long (https://github.com/hunterlong)"

RUN apk --no-cache add curl jq
ENV IS_DOCKER=true
ENV STATUP_DIR=/app

WORKDIR /app
VOLUME /app
EXPOSE 8080

HEALTHCHECK --interval=5s --timeout=5s --retries=5 CMD curl -s "http://localhost:8080/health" | jq -r -e ".online==true"

CMD ["statup"]
