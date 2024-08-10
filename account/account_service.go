package account

import (
	"command_line_clue/services/profile"
	"context"
)

type AccountService struct {
	profile.UnimplementedAccountServiceServer
}

func (a *AccountService) Login(ctx context.Context, profile *profile.Account) (*profile.Reply, error) {

}
func (a *AccountService) Logout(ctx context.Context, profile *profile.Account) (*profile.Reply, error) {

}
