package dao

import (
	"gorm.io/gorm"
	"via-blog/model"
	"via-blog/utils/errmsg"
)

// CreateArticle 新增文章
func CreateArticle(data *model.Article) int {
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCateArticles 查询指定分类下所有文章
func GetCateArticles(cid int, pageSize int, pageNum int) ([]model.Article, int, int64) {
	var cateArtList []model.Article
	var total int64

	err := DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid = ?", cid).Find(&cateArtList).Error
	DB.Model(&cateArtList).Where("cid = ?", cid).Count(&total)

	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

// GetArticle  查询单个文章
func GetArticle(id int) (model.Article, int) {
	var art model.Article
	err := DB.Where("id = ?", id).Preload("Category").First(&art).Error
	DB.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}



// GetArticles 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]model.Article, int, int64) {
	var artList []model.Article
	var total int64

	// 【表名是大小写敏感的】！！！！
	err := DB.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").Limit(
		pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Category").Find(&artList).Error

	// 单独计数
	DB.Model(&artList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return artList, errmsg.SUCCESS, total
}

// SearchArticle 搜索文章标题
func SearchArticle(title string, pageSize int, pageNum int) ([]model.Article, int, int64) {
	var artList []model.Article
	var total int64
	err := DB.Select("article.id, title, img, created_at, updated_at, `desc`, category.name, comment_count, read_count, category.name").Order(
		"created_at DESC").Joins("Category").Where("title LIKE ?", title+"%").Limit(
		pageSize).Offset((pageNum - 1) * pageSize).Find(&artList).Error

	DB.Model(&artList).Where("title LIKE ?", title + "%").Count(&total)

	if err != nil{
		return nil, errmsg.ERROR, 0
	}
	return artList, errmsg.SUCCESS, total
}

// UpdateArticle 编辑文章
func UpdateArticle(id int, data *model.Article) int {
	var art model.Article
	var maps = make(map[string]interface{})

	maps = map[string]interface{}{
		"title": data.Title,
		"cid": data.Cid,
		"desc": data.Desc,
		"content": data.Content,
		"img": data.Img,
	}

	err := DB.Model(&art).Where("id = ?", id).Updates(&maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var art model.Article
	err := DB.Where("id = ?", id).Delete(&art).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
