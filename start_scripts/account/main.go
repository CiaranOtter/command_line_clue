package main

import (
	"log"
	"net"

	database "github.com/CiaranOtter/command_line_clue/server/game_database"

	"github.com/CiaranOtter/command_line_clue/server/clc_services/profile"

	"github.com/CiaranOtter/command_line_clue/server/account_service"

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
