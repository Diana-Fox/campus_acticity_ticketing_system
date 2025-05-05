package server

import (
	context "context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/domain"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/service"
	api "ticket_utils/pkg/api/base"
	"ticket_utils/tool"
)

type GRPCImage struct {
	svc service.Image
	api.UnimplementedImageServiceServer
}

func (g *GRPCImage) Create(ctx context.Context, request *api.ImageListRequest) (*api.EmptyResponse, error) {
	listRequest := request.Items
	list := make([]*domain.Image, len(listRequest))
	for i := 0; i < len(listRequest); i++ {
		id, err := tool.NextId()
		if err != nil {
			return nil, err
		}
		l := &domain.Image{
			Category: int8(listRequest[i].Category),
			Id:       id,
			ImgUrl:   listRequest[i].ImgUrl,
			Sort:     int8(listRequest[i].Sort),
			TargetId: listRequest[i].TargetId,
			Type:     int8(listRequest[i].Type),
		}
		list = append(list, l)
	}
	err := g.svc.Create(ctx, list)
	return nil, err
}

func (g *GRPCImage) GetById(ctx context.Context, request *api.IdRequest) (*api.ImageListResponse, error) {
	list, err := g.svc.GetListByTargetId(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	images := make([]*api.Image, len(list))
	for i := 0; i < len(list); i++ {
		l := &api.Image{
			Id:       list[i].Id,
			ImgUrl:   list[i].ImgUrl,
			Sort:     int32(list[i].Sort),
			TargetId: list[i].TargetId,
			Type:     int32(list[i].Type),
			Category: int32(list[i].Category),
		}
		images = append(images, l)
	}
	return &api.ImageListResponse{
		Items: images,
	}, nil
}
