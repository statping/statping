FROM node:10.17.0 AS frontend
RUN npm install yarn -g
WORKDIR /statping
COPY ./frontend/package.json .
COPY ./frontend/yarn.lock .
RUN yarn install
COPY ./frontend .
RUN yarn build && rm -rf node_modules

# Statping Golang BACKEND building from source
# Creates "/go/bin/statping" and "/usr/local/bin/sass" for copying
FROM golang:1.14-alpine AS backend
LABEL maintainer="Hunter Long (https://github.com/hunterlong)"
ARG VERSION
RUN apk add --update --no-cache libstdc++ gcc g++ make git ca-certificates linux-headers wget curl jq
RUN curl -L -s https://assets.statping.com/sass -o /usr/local/bin/sass && \
    chmod +x /usr/local/bin/sass
WORKDIR /go/src/github.com/hunterlong/statping
ADD go.mod go.sum ./
RUN go mod download
ENV GO111MODULE on
RUN go get github.com/stretchr/testify/assert && \
    go get github.com/stretchr/testify/require && \
	go get github.com/GeertJohan/go.rice/rice && \
	go get github.com/cortesi/modd/cmd/modd && \
	go get github.com/crazy-max/xgo
COPY . .
COPY --from=frontend /statping/dist ./source/
RUN make clean generate embed build
RUN chmod a+x statping && mv statping /go/bin/statping

# Statping main Docker image that contains all required libraries
FROM alpine:latest
RUN apk --no-cache add libgcc libstdc++ curl jq

COPY --from=backend /go/bin/statping /usr/local/bin/
COPY --from=backend /usr/local/bin/sass /usr/local/bin/
COPY --from=backend /usr/local/share/ca-certificates /usr/local/share/

WORKDIR /app

ENV IS_DOCKER=true
ENV STATPING_DIR=/app
ENV PORT=8080

EXPOSE $PORT

HEALTHCHECK --interval=60s --timeout=10s --retries=3 CMD curl -s "http://localhost:$PORT/health" | jq -r -e ".online==true"

CMD statping -port $PORT

