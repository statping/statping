FROM alpine:latest
LABEL MAINTAINER = 'Hunter Long (https://github.com/hunterlong)'

# Locked version of Statup for 'latest' Docker tag
ENV VERSION=v0.43

RUN apk --no-cache add libstdc++ ca-certificates
RUN wget -q https://github.com/hunterlong/statup/releases/download/$VERSION/statup-linux-alpine.tar.gz && \
      tar -xvzf statup-linux-alpine.tar.gz && \
      chmod +x statup && \
      mv statup /usr/local/bin/statup

# sass Binary built for alpine linux
RUN wget -q https://assets.statup.io/sass && \
      chmod +x sass && \
      mv sass /usr/local/bin/sass

ENV IS_DOCKER=true
ENV SASS=/usr/local/bin/sass
ENV STATUP_DIR=/app

WORKDIR /app
VOLUME /app
EXPOSE 8080
ENTRYPOINT statup