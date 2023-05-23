FROM golang:1.19.2-alpine3.16

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY /migrator ./migrator
COPY internal/config ./internal/config
COPY migration ./migration
COPY .env .env

RUN CGO_ENABLED=0 go build -o /opt/migrator ./migrator

CMD [ "/opt/migrator", "-up=true"]