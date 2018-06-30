FROM alpine:latest

ENV VERSION=v0.27.81

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
ENV CMD_FILE=/usr/bin/cmd

RUN printf "#!/usr/bin/env sh\n\$1\n" > $CMD_FILE && \
      chmod +x $CMD_FILE

WORKDIR /app
#COPY build/statup-linux-alpine /usr/local/bin/statup

VOLUME /app

EXPOSE 8080
ENTRYPOINT statup