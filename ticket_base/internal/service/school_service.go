package service

import (
	"context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/model"
)

// SchoolService 学校表服务接口
type SchoolService interface {
	Create(ctx context.Context, school *model.School) error
	Update(ctx context.Context, school *model.School) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*model.School, error)
	List(ctx context.Context, page, pageSize int) ([]*model.School, error)
}
