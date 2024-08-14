package account_service

import (
	"clc_services/profile"
	"context"
	"database/sql"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountService struct {
	profile.UnimplementedAccountServiceServer
	DB *sql.DB
}

func (a *AccountService) Login(ctx context.Context, in *profile.Account) (*profile.Reply, error) {

	if in.GetRegister() {

		_, err := a.DB.Exec("INSERT INTO account (name, surnmane, usernamne) VALUES ('$1','$2','$3')", in.GetName(), in.GetSurname(), in.GetUsername())

		if err != nil {
			log.Fatalf("Failed to register new user", "error", err)
			return nil, err
		}

		return &profile.Reply{
			Success: true,
		}, nil
	}

	var id int = -1
	err := a.DB.QueryRow("SELECT id FROM account WHERE username=$1", in.GetUsername()).Scan(&id)

	if err != nil {
		log.Fatalf("Failed to register new user", "error", err)
		return nil, err
	}

	if id == -1 {
		return nil, status.Error(codes.NotFound, "Failed to find an account with the given username")
	}

	_, err = a.DB.Exec("UPDATE SET logged_in SET logged_in = true WHERE id = $1", id)

	if err != nil {
		log.Fatalf("Failed to register new user", "error", err)
		return nil, err
	}

	return &profile.Reply{
		Success: true,
	}, nil
}

func (a *AccountService) Logout(ctx context.Context, in *profile.Account) (*profile.Reply, error) {
	_, err := a.DB.Exec("UPDATE account SET logged_in = false WHERE usernmane = $1", in.GetUsername())
	if err != nil {
		return nil, status.Error(codes.NotFound, "No account with the given usernamen could be found")
	}

	return &profile.Reply{
		Success: true,
	}, nil
}
