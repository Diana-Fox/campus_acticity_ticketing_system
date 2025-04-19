package service

import (
	"campus_acticity_ticketing_system/ticket_base/internal/dao"
	"campus_acticity_ticketing_system/ticket_base/internal/model"
)

type TicketItemService interface {
	GetItemById(id int64) (*model.Item, error)
	GetItemCommentsByItemId(itemId int64) ([]*model.ItemComment, error)
	GetPlaceById(id int64) (*model.Place, error)
	GetPlaceSeatsByActivityId(activityId int64) ([]*model.PlaceSeat, error)
}

type ticketItemService struct {
	dao dao.TicketItemDAO
}

func NewTicketItemService(dao dao.TicketItemDAO) TicketItemService {
	return &ticketItemService{dao: dao}
}

func (s *ticketItemService) GetItemById(id int64) (*model.Item, error) {
	return s.dao.GetItemById(id)
}

func (s *ticketItemService) GetItemCommentsByItemId(itemId int64) ([]*model.ItemComment, error) {
	return s.dao.GetItemCommentsByItemId(itemId)
}

func (s *ticketItemService) GetPlaceById(id int64) (*model.Place, error) {
	return s.dao.GetPlaceById(id)
}

func (s *ticketItemService) GetPlaceSeatsByActivityId(activityId int64) ([]*model.PlaceSeat, error) {
	return s.dao.GetPlaceSeatsByActivityId(activityId)
}
