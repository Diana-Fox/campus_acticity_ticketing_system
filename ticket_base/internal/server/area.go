package server

import (
	context "context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/domain"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/service"
	api "ticket_utils/pkg/api/base"
	"ticket_utils/tool"
)

type GRPCArea struct {
	svc service.Area
	api.UnimplementedAreaServiceServer
}

func (g *GRPCArea) Create(ctx context.Context, request *api.AreaRequest) (*api.EmptyResponse, error) {
	id, err := tool.NextId()
	if err != nil {
		return nil, err
	}
	err = g.svc.Create(ctx, domain.Area{
		Id:       id,
		Name:     request.Name,
		ParentId: request.ParentId,
		Level:    int8(request.Level),
	})
	return nil, err
}

func (g *GRPCArea) Update(ctx context.Context, request *api.AreaRequest) (*api.EmptyResponse, error) {
	err := g.svc.Update(ctx, domain.Area{
		Id:       request.Id,
		Name:     request.Name,
		ParentId: request.ParentId,
		Level:    int8(request.Level),
	})
	return nil, err
}

func (g *GRPCArea) Delete(ctx context.Context, request *api.IdRequest) (*api.EmptyResponse, error) {
	err := g.svc.Delete(ctx, request.Id)
	return nil, err
}

func (g *GRPCArea) GetById(ctx context.Context, request *api.IdRequest) (*api.AreaResponse, error) {
	area, err := g.svc.GetById(ctx, request.Id)
	return &api.AreaResponse{
		Id:       area.Id,
		Name:     area.Name,
		ParentId: area.ParentId,
		Level:    int32(area.Level),
	}, err
}

func (g *GRPCArea) List(ctx context.Context) (*api.AreaListResponse, error) {
	list, err := g.svc.List(ctx)
	if err != nil {
		return nil, err
	}
	responses := make([]*api.AreaResponse, len(list))
	for i := 0; i < len(list); i++ {
		listResponse := &api.AreaResponse{
			Id:       list[i].Id,
			Name:     list[i].Name,
			ParentId: list[i].ParentId,
			Level:    int32(list[i].Level),
		}
		responses = append(responses, listResponse)
	}
	return &api.AreaListResponse{
		Items: responses,
	}, nil
}
