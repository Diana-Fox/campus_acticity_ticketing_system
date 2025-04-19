package service

import (
	"context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/model"
)

// ScopeService 作用域表服务接口
type ScopeService interface {
	Create(ctx context.Context, scope *model.Scope) error
	Update(ctx context.Context, scope *model.Scope) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*model.Scope, error)
	List(ctx context.Context, page, pageSize int) ([]*model.Scope, error)
}
