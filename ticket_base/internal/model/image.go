package model

import "time"

// Image 图片表
type Image struct {
	ID          int64     `json:"id"`
	ImgURL      string    `json:"img_url"`
	TargetID    int       `json:"target_id"`
	Sort        int       `json:"sort"`
	Type        int       `json:"type"`
	Category    int       `json:"category"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	Status      bool      `json:"status"`
}
