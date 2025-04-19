package domain

// model实体
type User struct {
	Id       uint64 `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Sex      int8   `json:"sex"`
}
