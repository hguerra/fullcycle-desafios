FROM golang:1.20 AS build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM hello-world:linux
WORKDIR /app
COPY --from=build /app/main ./
CMD [ "./main" ]
