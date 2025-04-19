package service

import (
	"context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base/internal/model"
)

type itemTypeServiceImpl struct {
	// 可以在这里添加数据访问层或其他依赖
}

func NewItemTypeService() ItemTypeService {
	return &itemTypeServiceImpl{}
}

func (s *itemTypeServiceImpl) Create(ctx context.Context, itemType *model.ItemType) error {
	// 实现创建逻辑
	return nil
}

func (s *itemTypeServiceImpl) Update(ctx context.Context, itemType *model.ItemType) error {
	// 实现更新逻辑
	return nil
}

func (s *itemTypeServiceImpl) Delete(ctx context.Context, id int64) error {
	// 实现删除逻辑
	return nil
}

func (s *itemTypeServiceImpl) GetByID(ctx context.Context, id int64) (*model.ItemType, error) {
	// 实现获取逻辑
	return nil, nil
}

func (s *itemTypeServiceImpl) List(ctx context.Context, page, pageSize int) ([]*model.ItemType, error) {
	// 实现列表逻辑
	return nil, nil
}
