FROM golang:1.11.1-alpine3.8 AS builder
RUN apk add --update bash make git
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/gobuffalo/packr/...
WORKDIR /go/src/github.com/vijayb8/crypto-api
COPY . ./
RUN make build

FROM alpine:3.8
ENV PORT_HTTPS=443
EXPOSE ${PORT_HTTPS}  
ENTRYPOINT ["/crypto-api"]