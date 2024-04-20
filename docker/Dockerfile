FROM golang:latest

WORKDIR /go/src/app

COPY . .

COPY .env .

RUN go mod download

RUN go build -o main ./cmd/main/main.go

EXPOSE 3000

CMD ["./main"]
