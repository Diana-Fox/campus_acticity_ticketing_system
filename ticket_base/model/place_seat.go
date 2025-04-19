package model

import "time"

type PlaceSeat struct {
	ID         int64     `json:"id"`
	X          int32     `json:"x"`
	Y          int32     `json:"y"`
	ActivityID int64     `json:"activity_id"`
	Sort       int32     `json:"sort"`
	State      int32     `json:"state"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	Status     bool      `json:"status"`
}