package main

import (
	"log"
	"net"

	"github.com/CiaranOtter/command_line_clue/server/clc_services/games"
	gameservice "github.com/CiaranOtter/command_line_clue/server/game_service"

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

	service := &gameservice.GameService{
		DB: database.OpenDB(),
	}

	games.RegisterGameServiceServer(server, service)

	err = server.Serve(list)

	if err != nil {
		panic(err)
	}

}
