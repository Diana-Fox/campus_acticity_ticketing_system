package service

import (
	"context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/model"
)

// ImageService 图片表服务接口
type ImageService interface {
	Create(ctx context.Context, image *model.Image) error
	Update(ctx context.Context, image *model.Image) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*model.Image, error)
	List(ctx context.Context, page, pageSize int) ([]*model.Image, error)
}
