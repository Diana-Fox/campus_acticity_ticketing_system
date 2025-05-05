package service

import (
	context "context"
	domain "github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_user/internal/internal/domian"
)

type user struct {
}

func (u *user) Register(ctx context.Context, d domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *user) LoginByPassword(ctx context.Context, d domain.User) string {
	//TODO implement me
	panic("implement me")
}

func (u *user) AddLinkUser(ctx context.Context, token string, names []string) error {
	//TODO implement me
	panic("implement me")
}

// user
func NewUser() User {
	return &user{}
}
