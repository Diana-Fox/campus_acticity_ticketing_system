package server

import (
	context "context"
	domain "github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_user/internal/internal/domian"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_user/internal/internal/service"
	api "ticket_utils/pkg/api/user"
	"ticket_utils/tool"
)

type GRPCUser struct {
	svc service.User
	api.UnimplementedUserServiceServer
}

func (g *GRPCUser) Register(ctx context.Context, request *api.RegisterRequest) (*api.RegisterResponse, error) {
	id, err := tool.NextId()
	if err != nil {
		return nil, err
	}
	err = g.svc.Register(ctx, domain.User{
		Id:       uint64(id),
		Phone:    request.Phone,
		Name:     request.Name,
		Email:    request.Email,
		IdCard:   request.CardId,
		Password: request.Password,
		Sex:      int8(request.Sex),
	})
	if err != nil {
		return nil, err
	}
	return &api.RegisterResponse{
		IsOk: true,
	}, nil
}

func (g *GRPCUser) LoginByPassword(ctx context.Context, request *api.LoginByPasswordRequest) (*api.UserResponse, error) {
	token := g.svc.LoginByPassword(ctx, domain.User{
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	})
	return &api.UserResponse{
		Token: token,
	}, nil
}

func (g *GRPCUser) AddLinkUser(ctx context.Context, request *api.AddLinkUserRequest) (*api.RegisterResponse, error) {
	err := g.svc.AddLinkUser(ctx, request.Token, request.Names)
	return &api.RegisterResponse{
		IsOk: true,
	}, err
}
