FROM alpine
LABEL maintainer="Hunter Long (https://github.com/hunterlong)"

ENV STATPING_VERSION=0.80.35

RUN apk add --no-cache ca-certificates linux-headers curl

RUN curl -L -s https://assets.statping.com/sass -o /usr/local/bin/sass && \
    chmod +x /usr/local/bin/sass

RUN curl -L -s https://github.com/hunterlong/statping/releases/download/v$STATPING_VERSION/statping-linux-alpine.tar.gz | tar -xz && \
    chmod +x statping && mv statping /usr/local/bin/statping

ENV DB_CONN=sqlite
ENV NAME="Statping Demo"
ENV DESCRIPTION="An Awesome Demo of a Statping Server running on Docker"
ENV DOMAIN=demo.statping.com
ENV SASS=/usr/local/bin/sass

ENV IS_DOCKER=true
ENV STATPING_DIR=/app
WORKDIR /app

COPY ./dev/demo-script.sh /app/

ENTRYPOINT ./demo-script.sh
