package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"viaBlog/model"
	"viaBlog/utils/errmsg"
)

// AddArticle 添加文字
func AddArticle(c *gin.Context)  {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code := model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCateArticles 查询指定分类下的所有文章
func GetCateArticles(c *gin.Context){
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cid, _ := strconv.Atoi(c.Param("id"))

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <0:
		pageSize = 10
	}

	if pageNum == 0{
		pageNum = 1
	}

	data, code, total := model.GetCateArticles(cid, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"total": total,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArticle 查询单个文章
func GetArticle(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := model.GetArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArticles 查询文章列表（模糊查询）
func GetArticles(c *gin.Context)  {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize < 0:
		pageSize = 10
	}

	if pageNum == 0{
		pageNum = 1
	}

	// todo 这里对title 的判断不应该放在 DAO 层吗？
	if len(title) ==0{
		data, code, total := model.GetArticles(pageSize, pageNum)
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data": data,
			"total": total,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	data, code, total := model.SearchArticle(title, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"total": total,
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateArticle 更新文章
func UpdateArticle(c *gin.Context)  {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.UpdateArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}