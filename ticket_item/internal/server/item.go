package server

import (
	context "context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_item/internal/domain"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_item/internal/service"
	api "ticket_utils/pkg/api/item"
	"ticket_utils/tool"
)

type GRPCItem struct {
	svc service.Item
	//api.UnimplementedItemServiceServer
}

func (g *GRPCItem) CreateItem(ctx context.Context, request *api.CreateItemRequest) (*api.EmptyResponse, error) {
	id, err := tool.NextId()
	if err != nil {
		return nil, err
	}
	err = g.svc.CreateItem(ctx, domain.Item{
		Id:                  id,
		ItemName:            request.Item.ItemName,
		AbstractMessage:     request.Item.AbstractMessage,
		ItemTypeId:          request.Item.ItemTypeId,
		ScopeId:             request.Item.ScopeId,
		BasicDescription:    request.Item.BasicDescription,
		ProjectDescription:  request.Item.ProjectDescription,
		ReminderDescription: request.Item.ReminderDescription,
		PlaceId:             request.Item.PlaceId,
		CommentCount:        0,
		StartTime:           request.Item.StartTime,
		EndTime:             request.Item.EndTime,
	})
	if err != nil {
		return nil, err
	}
	return &api.EmptyResponse{}, nil
}

func (g *GRPCItem) GetItem(ctx context.Context, request *api.GetItemRequest) (*api.EmptyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GRPCItem) UpdateItem(ctx context.Context, request *api.UpdateItemRequest) (*api.UpdateItemResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GRPCItem) DeleteItem(ctx context.Context, request *api.DeleteItemRequest) (*api.EmptyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GRPCItem) ListItem(ctx context.Context, request *api.ListItemRequest) (*api.ListItemResponse, error) {
	//TODO implement me
	panic("implement me")
}
