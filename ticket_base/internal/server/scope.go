package server

import (
	context "context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/domain"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/service"
	api "ticket_utils/pkg/api/base"
)

type GRPCScope struct {
	svc service.Scope
	api.UnimplementedScopeServiceServer
}

func (g *GRPCScope) Create(ctx context.Context, scope *api.Scope) (*api.EmptyResponse, error) {
	err := g.svc.Create(ctx, scope.Name)
	return &api.EmptyResponse{}, err
}

func (g *GRPCScope) Update(ctx context.Context, scope *api.Scope) (*api.EmptyResponse, error) {
	err := g.svc.Update(ctx, domain.Scope{
		Id:   scope.Id,
		Name: scope.Name,
	})
	return &api.EmptyResponse{}, err
}

func (g *GRPCScope) Delete(ctx context.Context, request *api.IdRequest) (*api.EmptyResponse, error) {
	err := g.svc.Delete(ctx, request.Id)
	return &api.EmptyResponse{}, err
}

func (g *GRPCScope) List(ctx context.Context, request *api.ListRequest) (*api.ScopeListResponse, error) {
	list, count, err := g.svc.List(ctx)
	if err != nil {
		return nil, err
	}
	scopes := make([]*api.Scope, len(list))
	for i := 0; i < len(scopes); i++ {
		l := &api.Scope{
			Id:   list[i].Id,
			Name: list[i].Name,
		}
		scopes[i] = l
	}
	return &api.ScopeListResponse{
		Items: scopes,
		Total: int32(count),
	}, nil
}
