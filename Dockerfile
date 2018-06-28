FROM alpine:latest

ENV VERSION=v0.27

RUN apk --no-cache add libstdc++ ca-certificates
RUN wget https://github.com/hunterlong/statup/releases/download/$VERSION/statup-linux-alpine.tar.gz && \
      tar -xvzf statup-linux-alpine.tar.gz && \
      chmod +x statup && \
      mv statup /usr/local/bin/statup

#COPY build/statup /usr/local/bin/statup
#RUN chmod +x /usr/local/bin/statup

RUN wget https://assets.statup.io/sass && \
      chmod +x sass && \
      mv sass /usr/local/bin/sass

ENV CMD_FILE=/usr/bin/cmd
RUN printf "#!/usr/bin/env sh\n\$1\n" > $CMD_FILE && \
      chmod +x $CMD_FILE
ENV USE_ASSETS=true

WORKDIR /app
VOLUME /app
EXPOSE 8080
ENTRYPOINT statup