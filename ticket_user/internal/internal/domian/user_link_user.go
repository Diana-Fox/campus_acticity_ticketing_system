package domain

// model实体
type UserLinkUser struct {
	Id     uint64 `json:"id"`
	UserId uint64 `json:"user_id"` //对应用户的主键
	Name   string `json:"name"`
	IdCard string `json:"id_card"` //身份证号
}
