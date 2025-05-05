package domain

// Place 场所
type Place struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Address   string `json:"address,omitempty"`
	AreaId    int64  `json:"area_id,omitempty"`
	AreaName  string `json:"area_name,omitempty"`
	XLength   int    `json:"x_length,omitempty"`
	YLength   int    `json:"y_length,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}
