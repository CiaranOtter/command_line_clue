OUT_DIR=clc_services

all:
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) protos/profile.proto
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) protos/message.proto
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) protos/games.proto
