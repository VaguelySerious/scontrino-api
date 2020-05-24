package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name  	 string `json:"name"`
	Password string `json:"-"`
}

type Group struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
}
