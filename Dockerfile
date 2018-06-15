FROM alpine

RUN apk add --no-cache libc6-compat

ENV VERSION="v0.12"

WORKDIR /app
RUN wget -q https://github.com/hunterlong/statup/releases/download/$VERSION/statup-linux-x64
RUN chmod +x statup-linux-x64 && mv statup-linux-x64 /usr/local/bin/statup

EXPOSE 8080

VOLUME /app

ENTRYPOINT statup