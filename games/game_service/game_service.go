package gameservice

import (
	"clc_services/games"
	"context"
	"database/sql"
)

type GameService struct {
	games.UnimplementedGameServiceServer
	DB *sql.DB
}

func (g GameService) RegisterNewGame(ctx context.Context, in *games.GameItem) (*games.Response, error) {
	name := in.GetName()
}
