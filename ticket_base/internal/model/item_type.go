package model

import "time"

// ItemType 商品类型表
type ItemType struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	Status      bool      `json:"status"`
}
