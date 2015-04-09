package entity

type User struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
