package domain

// model实体
type User struct {
	Id       uint64 `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IdCard   string `json:"id_card"` //身份证号
	Password string `json:"password"`
	Sex      int8   `json:"sex"`
}
