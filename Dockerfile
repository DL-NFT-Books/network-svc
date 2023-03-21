FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/dl-nft-books/network-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/network-svc /go/src/github.com/dl-nft-books/network-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/network-svc /usr/local/bin/network-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["network-svc"]
