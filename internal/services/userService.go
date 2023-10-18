package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"openim/internal/common"
	"openim/internal/dao"
)

type RegisterRequest struct {
	// Email    string `form:"email" json:"email" binding:"required"` //后面加上
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 用户注册接口
func UserRegister(c *gin.Context) {
	var register RegisterRequest
	if err := c.ShouldBind(&register); err != nil {
		common.BadRequest(c, "parameters error")
		return
	}

	// check same name
	n, err := dao.UserFindByName(register.UserName)
	if err != nil {
		log.Printf("find user by name %q error: %v", register.UserName, err)
		common.InternalServerError(c, "server error")
		return
	}
	fmt.Println("n:", n)
	if n > 0 {
		c.JSON(http.StatusConflict, gin.H{"status": "conflict", "message": "user name exists!"})
		return
	}

	// save data to database
	err = dao.UserRegister(register.UserName, register.Password)
	if err != nil {
		log.Println("register Error", err.Error())
		common.InternalServerError(c, "register error")
		return
	}

	common.OK(c, nil)
}
