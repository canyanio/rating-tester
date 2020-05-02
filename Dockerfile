FROM golang:1.13.5-alpine3.10 as builder
RUN apk update && apk upgrade && \
    apk add \
    xz-dev \
    musl-dev \
    gcc
RUN mkdir -p /go/src/github.com/canyanio/rating-tester
COPY . /go/src/github.com/canyanio/rating-tester
RUN cd /go/src/github.com/canyanio/rating-tester && env CGO_ENABLED=1 go build

FROM alpine:3.10
RUN apk update && apk upgrade && \
        apk add --no-cache ca-certificates xz
RUN mkdir -p /etc/rating-tester
COPY ./config.yaml /etc/rating-tester
COPY --from=builder /go/src/github.com/canyanio/rating-tester/rating-tester /usr/bin
ENTRYPOINT ["/usr/bin/rating-tester"]
