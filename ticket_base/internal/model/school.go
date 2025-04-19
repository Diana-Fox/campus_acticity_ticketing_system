package model

import "time"

// School 学校表
type School struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	Status      bool      `json:"status"`
}
