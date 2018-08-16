FROM alpine:latest

ENV VERSION=v0.41

RUN apk --no-cache add libstdc++ ca-certificates
RUN wget -q https://github.com/hunterlong/statup/releases/download/$VERSION/statup-linux-alpine.tar.gz && \
      tar -xvzf statup-linux-alpine.tar.gz && \
      chmod +x statup && \
      mv statup /usr/local/bin/statup

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