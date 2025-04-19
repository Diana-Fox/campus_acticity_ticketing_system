package model

import "time"

type ItemComment struct {
	ID         int64     `json:"id"`
	ItemID     int64     `json:"item_id"`
	UserID     int64     `json:"user_id"`
	Score      int32     `json:"score"`
	Content    string    `json:"content"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}