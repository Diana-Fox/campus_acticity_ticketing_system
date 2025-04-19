package dao

import (
	"campus_acticity_ticketing_system/ticket_base/internal/model"
	"database/sql"
)

type TicketItemDAO interface {
	GetItemById(id int64) (*model.Item, error)
	GetItemCommentsByItemId(itemId int64) ([]*model.ItemComment, error)
	GetPlaceById(id int64) (*model.Place, error)
	GetPlaceSeatsByActivityId(activityId int64) ([]*model.PlaceSeat, error)
}

type ticketItemDAO struct {
	db *sql.DB
}

func NewTicketItemDAO(db *sql.DB) TicketItemDAO {
	return &ticketItemDAO{db: db}
}

func (d *ticketItemDAO) GetItemById(id int64) (*model.Item, error) {
	// 实现获取Item的逻辑
}

func (d *ticketItemDAO) GetItemCommentsByItemId(itemId int64) ([]*model.ItemComment, error) {
	// 实现获取ItemComment的逻辑
}

func (d *ticketItemDAO) GetPlaceById(id int64) (*model.Place, error) {
	// 实现获取Place的逻辑
}

func (d *ticketItemDAO) GetPlaceSeatsByActivityId(activityId int64) ([]*model.PlaceSeat, error) {
	// 实现获取PlaceSeat的逻辑
}
