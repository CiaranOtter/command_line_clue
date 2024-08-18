package gameservice

import (
	"context"
	"database/sql"
	"log"

	"github.com/CiaranOtter/command_line_clue/server/clc_services/games"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GameService struct {
	games.UnimplementedGameServiceServer
	DB *sql.DB
}

func (g GameService) RegisterNewGame(ctx context.Context, in *games.GameItem) (*games.Response, error) {
	name := in.GetName()

	_, err := g.DB.Exec("INSERT INTO registered_games (name) VALUES ($1)", name)

	if err != nil {
		log.Printf(err.Error())
		return nil, status.Error(codes.Canceled, "Failed to saved game into list of registered games")
	}

	return &games.Response{
		Success: true,
	}, nil
}
