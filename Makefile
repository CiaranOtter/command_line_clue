OUT_DIR=services

all:
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) protos/profile.proto

clean:
	rm -rf services/*.proto