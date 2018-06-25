FROM alpine:latest

ENV VERSION=v0.22

RUN apk --no-cache add ca-certificates
RUN wget https://github.com/hunterlong/statup/releases/download/$VERSION/statup-linux-alpine && \
      chmod +x statup-linux-alpine && \
      mv statup-linux-alpine /usr/local/bin/statup
WORKDIR /app
VOLUME /app
RUN statup version
EXPOSE 8080
ENTRYPOINT statup