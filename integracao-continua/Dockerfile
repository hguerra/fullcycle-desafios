FROM golang:1.21.3-alpine

WORKDIR /app

COPY . .

RUN go build -o ./bin/math main.go

CMD ["./bin/math"]
