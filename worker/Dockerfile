FROM golang:1.21

WORKDIR /app
COPY ../proto ./proto
COPY . .

RUN go mod init worker && go get github.com/streadway/amqp && go get ./...
RUN go build -o worker main.go

CMD ["./worker"]
