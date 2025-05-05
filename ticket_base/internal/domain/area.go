package domain

// Area 区域
type Area struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	ParentId int64  `json:"parent_id,omitempty"`
	Level    int8   `json:"level,omitempty"`
	Children []Area `json:"children,omitempty"`
}
