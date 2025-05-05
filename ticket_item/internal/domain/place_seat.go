package domain

// PlaceSeat 场所座位
type PlaceSeat struct {
	Id         int64 `json:"id,omitempty"`
	X          int64 `json:"x,omitempty"`
	Y          int64 `json:"y,omitempty"`
	ActivityId int64 `json:"activityId,omitempty"`
	Status     int8  `json:"status,omitempty"`
}
