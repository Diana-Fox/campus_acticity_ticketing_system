package domain

// Image 图的相关字段
type Image struct {
	Id       int64  `json:"id,omitempty"`
	ImgUrl   string `json:"img_url,omitempty"`
	TargetId int64  `json:"target_id,omitempty"`
	Sort     int8   `json:"sort,omitempty"`
	Type     int8   `json:"type,omitempty"`
	Category int8   `json:"category,omitempty"`
}
