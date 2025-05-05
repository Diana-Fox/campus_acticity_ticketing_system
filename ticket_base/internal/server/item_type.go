package server

import (
	context "context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/domain"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/service"
	api "ticket_utils/pkg/api/base"
)

type GRPCItemType struct {
	svc service.ItemType
	api.UnimplementedItemTypeServiceServer
}

func (g *GRPCItemType) Create(ctx context.Context, itemType *api.ItemType) (*api.EmptyResponse, error) {
	err := g.svc.Create(ctx, itemType.Name)
	return &api.EmptyResponse{}, err
}

func (g *GRPCItemType) Update(ctx context.Context, itemType *api.ItemType) (*api.EmptyResponse, error) {
	err := g.svc.Update(ctx, domain.ItemType{
		Id:   itemType.Id,
		Name: itemType.Name,
	})
	return &api.EmptyResponse{}, err
}

func (g *GRPCItemType) Delete(ctx context.Context, request *api.IdRequest) (*api.EmptyResponse, error) {
	err := g.svc.Delete(ctx, request.Id)
	return &api.EmptyResponse{}, err
}

//func (g *GRPCItemType) GetById(ctx context.Context, request *api.IdRequest) (*api.ItemType, error) {
//	item, err := g.svc.GetById(ctx, request.Id)
//	if err != nil {
//		return nil, err
//	}
//	return &api.ItemType{
//		Id:   item.Id,
//		Name: item.Name,
//	}, err
//}

func (g *GRPCItemType) List(ctx context.Context, request *api.ListRequest) (*api.ItemTypeListResponse, error) {
	list, count, err := g.svc.List(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*api.ItemType, len(list))
	for i := 0; i < len(list); i++ {
		l := &api.ItemType{
			Id:   list[i].Id,
			Name: list[i].Name,
		}
		items[i] = l
	}
	return &api.ItemTypeListResponse{
		Items: items,
		Total: int32(count),
	}, err
}
