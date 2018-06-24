FROM alpine:latest

ENV VERSION=v0.22

RUN apk --no-cache add ca-certificates
RUN wget https://github.com/hunterlong/statup/releases/download/$VERSION/statup-alpine && \
      chmod +x statup-alpine && \
      mv statup-alpine /usr/local/bin/statup
WORKDIR /app
VOLUME /app
RUN statup version
EXPOSE 8080
ENTRYPOINT statup