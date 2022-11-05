package dao

import (
	"gorm.io/gorm"
	"via-blog/model"
	"via-blog/utils/errmsg"
)

// CheckCategory 查看分类是否存在
func CheckCategory(name string) int {
	var cate model.Category
	DB.Select("id").Where("name = ?", name).Find(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// CreateCategory 新增分类
func CreateCategory(data *model.Category) int {
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCategory 查询单个分类信息
func GetCategory(id int) (model.Category, int) {
	var cate model.Category
	DB.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCESS
}

// GetCategories 查询分类列表
func GetCategories(pageSize int, pageNum int) ([]model.Category, int64) { // 第二个参数是结果总数
	var cates []model.Category
	var total int64 // 结果总数
	err = DB.Find(&cates).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	DB.Model(&cates).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cates, total
}

// UpdateCategory 编辑（更新）分类信息
func UpdateCategory(id int, data *model.Category) int {
	var cate model.Category
	var maps = make(map[string]interface{})

	maps["name"] = data.Name
	err := DB.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

// DeleteCategory 删除分类
func DeleteCategory(id int)  int{
	var cate model.Category
	err := DB.Where("id = ?", id).Delete(&cate).Error
	if err!= nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
