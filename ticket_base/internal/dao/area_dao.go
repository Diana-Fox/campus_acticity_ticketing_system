package dao

import (
	"context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/model"
)

// AreaDAO 区域表数据访问接口
type AreaDAO interface {
	Create(ctx context.Context, area *model.Area) error
	Update(ctx context.Context, area *model.Area) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*model.Area, error)
	List(ctx context.Context, page, pageSize int) ([]*model.Area, error)
}
