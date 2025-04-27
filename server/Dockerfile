FROM golang:1.21

WORKDIR /app
COPY ../proto ./proto
COPY . .

RUN go mod init server && go get google.golang.org/grpc github.com/streadway/amqp && go get ./...
RUN go build -o server main.go

CMD ["./server"]
