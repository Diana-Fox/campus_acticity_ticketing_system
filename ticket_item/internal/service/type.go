package service

import (
	context "context"
	"github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_item/internal/domain"
)

type Item interface {
	CreateItem(ctx context.Context, item domain.Item) error
	GetItem(ctx context.Context, id int64) (domain.Item, error)
}
type ItemComment interface {
}
type Place interface {
}
type PlaceItem interface {
}
