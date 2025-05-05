package service

import (
	context "context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/domain"
)

type Area interface {
	Create(ctx context.Context, area domain.Area) error
	GetById(ctx context.Context, id int64) (domain.Area, error)
	List(ctx context.Context) ([]domain.Area, error)
	Update(ctx context.Context, area domain.Area) error
	Delete(ctx context.Context, id int64) error
}

type Image interface {
	Create(ctx context.Context, list []*domain.Image) error
	GetListByTargetId(ctx context.Context, id int64) ([]*domain.Image, error) //一般批量存图
}

type ItemType interface {
	Create(ctx context.Context, name string) error
	Update(ctx context.Context, itemType domain.ItemType) error
	Delete(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (domain.ItemType, error)
	List(ctx context.Context) ([]domain.ItemType, int64, error)
}
type School interface {
	Create(ctx context.Context, d *domain.School) error
	Update(ctx context.Context, d *domain.School) error
	Delete(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (*domain.School, error)
	List(ctx context.Context, page int32, size int32, name string) ([]*domain.School, int64, error)
}
type Scope interface {
	Create(ctx context.Context, name string) error
	Update(ctx context.Context, scope domain.Scope) error
	Delete(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (domain.Scope, error)
	List(ctx context.Context) ([]domain.Scope, int64, error)
}
