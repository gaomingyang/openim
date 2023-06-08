package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"openim/dao"
)

type RegisterRequest struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginRequest struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// register
func UserRegister(c *gin.Context) {
	var register RegisterRequest
	if err := c.ShouldBind(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "parameters error"})
		return
	}

	// check same name
	n, err := dao.UserFindByName(register.UserName)
	fmt.Println("n:", n)
	if n > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "repeat user name"})
		return
	}

	// save data to database
	err = dao.UserRegister(register.UserName, register.Password)
	if err != nil {
		log.Println("register Error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "register error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// login
func UserLogin(c *gin.Context) {
	var login LoginRequest
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "parameters error"})
		return
	}

	user, err := dao.UserLogin(login.UserName, login.Password)
	log.Println("user:")
	log.Printf("%+v\n", user)
	if err != nil {
		log.Println("login Error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "login error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
}
