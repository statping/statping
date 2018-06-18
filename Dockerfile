FROM alpine

#RUN apk add --no-cache libc6-compat

ENV VERSION="v0.14"
#RUN wget -q https://github.com/hunterlong/statup/releases/download/$VERSION/statup-linux-static
#RUN chmod +x statup-linux-static && mv statup-linux-static /usr/local/bin/statup

EXPOSE 8080

VOLUME /app

ENTRYPOINT statup