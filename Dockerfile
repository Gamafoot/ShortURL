FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

RUN go build -o /app/build/run /app/cmd/main.go

EXPOSE 8000

CMD ["/app/build/run"]