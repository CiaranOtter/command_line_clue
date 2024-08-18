package main

import (
	"log"
	"net"

	"github.com/CiaranOtter/command_line_clue/server/clc_services/message"
	"github.com/CiaranOtter/command_line_clue/server/message_service"

	database "github.com/CiaranOtter/command_line_clue/server/game_database"

	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	defer list.Close()

	server := grpc.NewServer()
	defer server.GracefulStop()

	service := &message_service.MessageServer{
		DB:           database.OpenDB(),
		Sender_chans: make(map[string]chan *message.ReceiveMessage),
	}

	message.RegisterMessageServiceServer(server, service)

	err = server.Serve(list)

	if err != nil {
		panic(err)
	}

}
