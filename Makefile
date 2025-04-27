PROTO_DIR=proto
PROTO_FILE=$(PROTO_DIR)/homework.proto
GEN_GO_DIR=proto

all: deps proto

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

proto:
	protoc --go_out=$(GEN_GO_DIR) --go-grpc_out=$(GEN_GO_DIR) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_FILE)

clean:
	rm -f $(GEN_GO_DIR)/*.pb.go
