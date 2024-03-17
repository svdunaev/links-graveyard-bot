FROM golang:1.22-alpine

RUN apk update \
    && apk add --no-cache git \
    && apk add --update gcc musl-dev

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

ENV CGO_ENABLED=1
RUN go build -o /links-graveyard

EXPOSE 8080

CMD ["/links-graveyard"]