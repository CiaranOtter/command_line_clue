all:
	protoc --go_out=. --go-grpc_out=. proto/clue.proto