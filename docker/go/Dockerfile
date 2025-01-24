FROM golang:1.20.2-alpine

COPY src/api /go/src/api/

WORKDIR /go/src/api/

COPY src/api/go.mod /go/src/api/
COPY src/api/go.sum /go/src/api/
RUN go mod download

RUN apk update \
&& apk add --no-cache git

EXPOSE 8080
