package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gambar   string `json:"gambar"`
}

type UserDetail struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Gambar   string `json:"gambar"`
}
