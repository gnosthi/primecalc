FROM golang:1.10-alpine AS builder

RUN apk add -U make gcc musl-dev ncurses git

ADD . /go/src/github.com/gnosthi/primecalc
WORKDIR /go/src/github.com/gnosthi/primecalc

RUN TERM=vt100 make

FROM alpine:3.7

COPY --from=0 /go/src/github.com/gnosthi/primecalc/primecalc /usr/bin/primecalc

ENTRYPOINT ["/usr/bin/primecalc"]
