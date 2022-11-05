package dao

import (
	"via-blog/model"
	"via-blog/utils/errmsg"
)

// GetProfile 获取个人信息设置
func GetProfile(id int) (model.Profile, int) {
	var profile model.Profile
	err = DB.Where("ID = ?", id).First(&profile).Error
	if err !=nil{
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS

}

// UpdateProfile 更新个人信息蛇追
func UpdateProfile(id int, data *model.Profile) int {
	var profile model.Profile
	err := DB.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS


}