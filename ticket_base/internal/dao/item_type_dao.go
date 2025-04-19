package dao

import (
	"context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/model"
)

// ItemTypeDAO 商品类型表数据访问接口
type ItemTypeDAO interface {
	Create(ctx context.Context, itemType *model.ItemType) error
	Update(ctx context.Context, itemType *model.ItemType) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*model.ItemType, error)
	List(ctx context.Context, page, pageSize int) ([]*model.ItemType, error)
}
