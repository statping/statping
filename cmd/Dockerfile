FROM alpine:latest
LABEL MAINTAINER = 'Hunter Long (https://github.com/hunterlong)'

# Statup 'dev' image that will download the latest version of Statup from github

RUN apk --no-cache add libstdc++ ca-certificates curl grep sed

RUN export VERSION=$(curl --silent "https://api.github.com/repos/hunterlong/statup/releases/latest" \
    | grep '"tag_name":' \
    | sed -E 's/.*"([^"]+)".*/\1/'); \
    wget -q https://github.com/hunterlong/statup/releases/download/$VERSION/statup-linux-alpine.tar.gz;

RUN tar -xvzf statup-linux-alpine.tar.gz && \
    chmod +x statup && \
    mv statup /usr/local/bin/statup

RUN statup version

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