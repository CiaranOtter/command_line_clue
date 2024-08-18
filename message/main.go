package main

import (
	"clc_services/message"
	database "game_database"
	"log"
	"message_service/message_service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatal(err)
	}

	defer list.Close()

	server := grpc.NewServer()
	defer server.GracefulStop()

	service := &message_service.MessageServer{
		DB:           database.OpenDB(),
		Sender_chans: make(map[string](chan *message.ReceiveMessage)),
	}

	message.RegisterMessageServiceServer(server, service)

	err = server.Serve(list)

	if err != nil {
		panic(err)
	}

}
