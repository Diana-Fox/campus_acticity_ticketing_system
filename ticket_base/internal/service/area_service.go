package service

import (
	"context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/model"
)

// AreaService 区域表服务接口
type AreaService interface {
	Create(ctx context.Context, area *model.Area) error
	Update(ctx context.Context, area *model.Area) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*model.Area, error)
	List(ctx context.Context, page, pageSize int) ([]*model.Area, error)
}
