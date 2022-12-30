FROM golang:1.19-bullseye

EXPOSE 80
WORKDIR /app

COPY ./src ./src
RUN go build -a -x -o main ./src/server.go

CMD ["./main"]
