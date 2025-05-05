package domain

// School 学校
type School struct {
	Id        int64   `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Code      string  `json:"code,omitempty"`      //注册码
	Url       string  `json:"url"`                 //验证校园的接口链接
	Latitude  float32 `json:"latitude,omitempty"`  //维度
	Longitude float32 `json:"longitude,omitempty"` //经度
	Address   string  `json:"address,omitempty"`   //地址信息
}
