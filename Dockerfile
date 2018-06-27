FROM alpine:latest

ENV VERSION=v0.25

RUN apk --no-cache add ca-certificates
RUN wget https://github.com/hunterlong/statup/releases/download/$VERSION/statup-linux-alpine.tar.gz && \
      tar -xvzf statup-linux-alpine.tar.gz && \
      chmod +x statup && \
      mv statup /usr/local/bin/statup
WORKDIR /app
VOLUME /app
RUN statup version
EXPOSE 8080
ENTRYPOINT statup