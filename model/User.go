package model

import (
	bcrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;unique" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2"`
}

// BeforeCreate 固定方法
func (this *User) BeforeCreate (_ *gorm.DB) (err error)  {
	this.Password = ScryptPw(this.Password)
	return nil
}

// BeforeUpdate 固定方法
func (this *User) BeforeUpdate(_ *gorm.DB) (err error) {
	this.Password = ScryptPw(this.Password)
	return nil
}



// ScryptPw 生成密码
func ScryptPw(password string)  string {
	const cost = 10
	HashPW, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPW)
}

