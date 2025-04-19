package dao

import (
	"context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/model"
)

// ScopeDAO 作用域表数据访问接口
type ScopeDAO interface {
	Create(ctx context.Context, scope *model.Scope) error
	Update(ctx context.Context, scope *model.Scope) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*model.Scope, error)
	List(ctx context.Context, page, pageSize int) ([]*model.Scope, error)
}
