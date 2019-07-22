FROM golang:1.11.1-alpine3.8 AS builder
RUN apk add --update bash make git
RUN go get "github.com/go-chi/chi"
RUN	go get "github.com/go-chi/cors"
RUN	go get "github.com/kelseyhightower/envconfig"
RUN	go get "github.com/sirupsen/logrus"
WORKDIR /go/src/github.com/vijayb8/crypto-api
COPY . ./
RUN make build

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/github.com/vijayb8/crypto-api/build/crypto-api /
ENV PORT_HTTPS=8080
EXPOSE ${PORT_HTTPS}  
ENTRYPOINT ["/crypto-api"]