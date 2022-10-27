package model

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120"`
	Role int `gorm:"type:int;DEFAULT:2" json:"role" validate:"required, gte=2"`
}


