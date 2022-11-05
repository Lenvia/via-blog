package dao

import (
	"golang.org/x/crypto/bcrypt"
	"via-blog/model"
	"via-blog/utils/errmsg"
)

// CheckLogin 后台登录验证
func CheckLogin(username string, password string) (model.User, int) {
	var user model.User
	var PasswordErr error

	DB.Where("username = ?", username).First(&user)
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0{  // 用户不存在
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil{  // 密码不正确
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1{  // 非管理员
		return user, errmsg.ERROR_USER_NO_RIGHT
	}

	return user, errmsg.SUCCESS
}

// CheckLoginFront todo 前台登录验证
func CheckLoginFront(username string, password string) (model.User, int) {
	var user model.User
	var PasswordErr error

	DB.Where("username = ?", username).First(&user)
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0{  // 用户不存在
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil{  // 密码不正确
		return user, errmsg.ERROR_PASSWORD_WRONG
	}

	return user, errmsg.SUCCESS
}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var user model.User
	DB.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 { // 已有数据
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CheckUpdateUser 检查将要更新的用户名是否已经存在
func CheckUpdateUser(id int, name string) int {
	var user model.User
	DB.Select("id = ?", id).Where("username = ?", name).First(&user)
	if user.ID == uint(id){
		return errmsg.SUCCESS
	}
	if user.ID > 0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *model.User) int {
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUser 查询单个用户
func GetUser(id int) (model.User, int) {
	var user model.User
	err := DB.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil{
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]model.User, int64){  // 第二个参数是结果总数
	var users []model.User
	var total int64

	if username != ""{  // 模糊查询
		DB.Select("id, username, role, created_at").Where(
			"username LIKE ?", username+ "%",).Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users)

		DB.Model(&users).Where("username LIKE ?", username + "%").Count(&total)
		return users, total
	}
	// 否则查找全部，并且只返回部分字段
	err = DB.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error
	DB.Model(&users).Count(&total)

	if err != nil{
		return users, 0
	}
	return users, total
}

// UpdateUser 更新用户
func UpdateUser(id int, data * model.User) int {
	var user model.User
	var maps = make(map[string]interface{})

	// todo 我觉得这里可以用反射来确定map的键
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := DB.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ChangePassWord 修改密码
func ChangePassWord(id int, data *model.User) int {
	err := DB.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user model.User
	err := DB.Where("id = ?", id).Delete(&user).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}