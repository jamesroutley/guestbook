FROM golang:1.11-alpine as builder
RUN apk add --no-cache gcc musl-dev
RUN mkdir /build
ADD . /go/src/github.com/jamesroutley/guestbook/guestbook-server
WORKDIR /go/src/github.com/jamesroutley/guestbook/guestbook-server
RUN go build .

FROM alpine
RUN apk update \
    && apk add sqlite
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /go/src/github.com/jamesroutley/guestbook/guestbook-server /app/
WORKDIR /app
CMD ["./guestbook-server"]
