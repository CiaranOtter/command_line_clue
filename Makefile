OUT_DIR = /Users/ciaranotter/Documents/personal/command_line_clue/web_app/client/src/comm

all:
	protoc --go_out=. --go-grpc_out=. proto/clue.proto
