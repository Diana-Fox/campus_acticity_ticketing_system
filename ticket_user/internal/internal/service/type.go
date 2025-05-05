package service

import (
	context "context"
	domain "github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_user/internal/internal/domian"
)

type User interface {
	Register(ctx context.Context, d domain.User) error
	LoginByPassword(ctx context.Context, d domain.User) string
	AddLinkUser(ctx context.Context, token string, names []string) error
}
