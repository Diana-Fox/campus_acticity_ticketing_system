package server

import (
	context "context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/domain"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/service"
	api "ticket_utils/pkg/api/base"
	"ticket_utils/tool"
)

// 学校信息
type GRPCSchool struct {
	svc service.School
	api.UnimplementedSchoolServiceServer
}

func (g *GRPCSchool) Create(ctx context.Context, school *api.School) (*api.EmptyResponse, error) {
	id, err := tool.NextId()
	if err != nil {
		return nil, err
	}
	err = g.svc.Create(ctx, &domain.School{
		Id:        id,
		Code:      school.Code,
		Url:       school.Url,
		Latitude:  school.Latitude,
		Longitude: school.Longitude,
		Name:      school.Name,
	})
	return &api.EmptyResponse{}, err
}

func (g *GRPCSchool) Update(ctx context.Context, school *api.School) (*api.EmptyResponse, error) {
	err := g.svc.Update(ctx, &domain.School{
		Id:        school.Id,
		Code:      school.Code,
		Latitude:  school.Latitude,
		Longitude: school.Longitude,
		Name:      school.Name,
	})
	return &api.EmptyResponse{}, err
}

func (g *GRPCSchool) Delete(ctx context.Context, request *api.IdRequest) (*api.EmptyResponse, error) {
	err := g.svc.Delete(ctx, request.Id)
	return &api.EmptyResponse{}, err
}

func (g *GRPCSchool) GetById(ctx context.Context, request *api.IdRequest) (*api.School, error) {
	school, err := g.svc.GetById(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &api.School{
		Id:      school.Id,
		Name:    school.Name,
		Code:    school.Code,
		Url:     school.Url,
		Address: school.Address,
	}, nil
}

func (g *GRPCSchool) List(ctx context.Context, request *api.SchoolListRequest) (*api.SchoolListResponse, error) {
	list, count, err := g.svc.List(ctx, request.Page, request.PageSize, request.Name)
	if err != nil {
		return nil, err
	}
	schools := make([]*api.School, len(list))
	for i := 0; i < len(schools); i++ {
		l := &api.School{
			Id:      list[i].Id,
			Name:    list[i].Name,
			Code:    list[i].Code,
			Url:     list[i].Url,
			Address: list[i].Address,
		}
		schools[i] = l
	}
	return &api.SchoolListResponse{
		Items: schools,
		Total: count,
	}, nil
}
