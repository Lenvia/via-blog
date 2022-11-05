package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"via-blog/dao"
	"via-blog/model"
	"via-blog/utils/errmsg"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)

	code := dao.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		dao.CreateCategory(&data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetCategory 查询分类信息
func GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := dao.GetCategory(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetCategories 查询分类列表
func GetCategories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize < 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := dao.GetCategories(pageSize, pageNum)
	code := errmsg.SUCCESS

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// UpdateCategory 更新分类
func UpdateCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := dao.CheckCategory(data.Name)  // 查询要更新的类别名是否已经存在
	if code == errmsg.SUCCESS{
		dao.UpdateCategory(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED{
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := dao.DeleteCategory(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
