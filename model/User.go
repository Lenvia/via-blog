package model

import (
	bcrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"via-blog/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;unique" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required, gte=2"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 { // 已有数据
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CheckUpdateUser 检查将要更新的用户名是否已经存在
func CheckUpdateUser(id int, name string) int {
	var user User
	db.Select("id = ?", id).Where("username = ?", name).First(&user)
	if user.ID == uint(id){
		return errmsg.SUCCESS
	}
	if user.ID > 0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUser 查询单个用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil{
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64){  // 第二个参数是结果总数
	var users []User
	var total int64

	if username != ""{  // 模糊查询
		db.Select("id, username, role, created_at").Where(
			"username LIKE ?", username+ "%",).Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users)

		db.Model(&users).Where("username LIKE ?", username + "%").Count(&total)
		return users, total
	}
	// 否则查找全部，并且只返回部分字段
	err = db.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error
	db.Model(&users).Count(&total)

	if err!= nil{
		return users, 0
	}
	return users, total
}

// UpdateUser 更新用户
func UpdateUser(id int, data * User) int {
	var user User
	var maps = make(map[string]interface{})

	// todo 我觉得这里可以用反射来确定map的键
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ChangePassWord 修改密码
func ChangePassWord(id int, data *User) int {
	err := db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
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

// CheckLogin todo 后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	//var PasswordErr error

	return user, errmsg.SUCCESS
}

// CheckLoginFront todo 前台登录验证
func CheckLoginFront(username string, password string) (User, int) {
	var user User
	//var PasswordErr error

	return user, errmsg.SUCCESS
}