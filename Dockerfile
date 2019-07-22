FROM golang:1.11.1-alpine3.8 AS builder
RUN apk add --update bash make git
RUN go get "github.com/go-chi/chi"
RUN	go get "github.com/go-chi/cors"
RUN	go get "github.com/kelseyhightower/envconfig"
RUN	go get "github.com/sirupsen/logrus"
WORKDIR /go/src/github.com/vijayb8/crypto-api
COPY . ./
RUN make build

FROM alpine:3.8
COPY --from=builder /go/src/github.com/vijayb8/crypto-api /
ENV PORT_HTTPS=443
EXPOSE ${PORT_HTTPS}  
ENTRYPOINT ["build/crypto-api"]