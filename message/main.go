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
	list, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	defer list.Close()

	server := grpc.NewServer()
	defer server.GracefulStop()

	service := &message_service.MessageServer{
		DB: database.OpenDB(),
	}

	message.RegisterMessageServiceServer(server, service)

	err = server.Serve(list)

	if err != nil {
		panic(err)
	}

}
