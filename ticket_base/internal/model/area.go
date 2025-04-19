package model

import "time"

// Area 区域表
type Area struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Parent      string    `json:"parent"`
	Level       int       `json:"level"`
	IsHot       int       `json:"is_hot"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	Status      bool      `json:"status"`
}
