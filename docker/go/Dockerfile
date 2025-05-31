FROM golang:1.23-alpine

COPY src/api /go/src/api/

WORKDIR /go/src/api/

RUN go install github.com/air-verse/air@latest

COPY src/api/go.mod /go/src/api/
COPY src/api/go.sum /go/src/api/

RUN apk update \
&& apk add --no-cache git

CMD ["air", "-c", ".air.toml"]

EXPOSE 8080
