FROM golang:1.10.3 as builder
WORKDIR /go/src/github.com/hunterlong/statup
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -o statup .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/hunterlong/statup /app/
RUN chmod +x /app/statup
CMD ["./statup", "version"]