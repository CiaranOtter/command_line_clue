package main

import (
	"account_service/account_service"
	"clc_services/profile"
	database "game_database"
	"log"
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

	service := &account_service.AccountService{
		DB: database.OpenDB(),
	}

	profile.RegisterAccountServiceServer(server, service)

	err = server.Serve(list)

	if err != nil {
		panic(err)
	}

}
