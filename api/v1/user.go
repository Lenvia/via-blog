package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"viaBlog/model"
	"viaBlog/utils/errmsg"
	"viaBlog/utils/validator"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)  // 验证输入是否合法（不查询数据库）
	if validCode != errmsg.SUCCESS{
		c.JSON(
			http.StatusOK, gin.H{
				"status" : validCode,
				"message" : msg,
			},
		)
		c.Abort()
		return
	}

	code := model.CheckUser(data.Username) // 查看用户名是否已经存在
	if code == errmsg.SUCCESS{
		model.CreateUser(&data)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"message" : errmsg.GetErrMsg(code),
		},
	)
}

// GetUser 查询单个用户
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})

	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"data": maps,
			"total": 1,
			"message" : errmsg.GetErrMsg(code),
		},
	)
}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")  // 如果为空则查询全部，不为空则进行模糊查询

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize < 0:
		pageSize = 0
	}

	if pageNum == 0{
		pageNum = 1
	}

	data, total := model.GetUsers(username, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"data": data,
			"total": total,
			"message" : errmsg.GetErrMsg(code),
		},
	)
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context)  {
	var data model.User
	id, _ :=strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&data)

	code := model.CheckUpdateUser(id, data.Username)  // 检查将要更新的用户名是否已经存在
	if code == errmsg.SUCCESS{
		model.UpdateUser(id, &data)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"message" : errmsg.GetErrMsg(code),
		},
	)

}

// ChangeUserPassWord 修改密码
func ChangeUserPassWord(c *gin.Context)  {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.ChangePassWord(id, &data)
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"message" : errmsg.GetErrMsg(code),
		},
	)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"message" : errmsg.GetErrMsg(code),
		},
	)
}
