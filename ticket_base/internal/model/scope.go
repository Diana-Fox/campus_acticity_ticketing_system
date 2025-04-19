package model

// Scope 作用域表
type Scope struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
