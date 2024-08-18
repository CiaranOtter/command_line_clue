package gameservice

import (
	"context"
	"database/sql"

	"github.com/CiaranOtter/command_line_clue/server/clc_services/games"
)

type GameService struct {
	games.UnimplementedGameServiceServer
	DB *sql.DB
}

func (g GameService) RegisterNewGame(ctx context.Context, in *games.GameItem) (*games.Response, error) {
	name := in.GetName()

	g.DB.Exec("INSERT")
}
