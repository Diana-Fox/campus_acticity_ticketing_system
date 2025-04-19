package model

import "time"

type Place struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	AreaID     int64     `json:"area_id"`
	AreaName   string    `json:"area_name"`
	XLength    int32     `json:"x_length"`
	YLength    int32     `json:"y_length"`
	Latitude   string    `json:"latitude"`
	Longitude  string    `json:"longitude"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	Status     bool      `json:"status"`
}