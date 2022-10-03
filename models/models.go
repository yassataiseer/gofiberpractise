package models

import "gorm.io/gorm"

type Notes struct {
		gorm.Model

		Title string
		Text string
		Id uint
		User string `gorm:"primaryKey;autoIncrement:true"`
}

type User struct {
	gorm.Model

	Username string
	Passsword string
}
