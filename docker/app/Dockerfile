FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o /opt/app ./main.go

CMD ["/opt/app"]