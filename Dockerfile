FROM golang:alpine

WORKDIR /app

COPY main.go go.mod ./
RUN go build -o app

EXPOSE 8080

CMD ["./app"]