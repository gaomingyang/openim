package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

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

type LoginResponse struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
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
	if err != nil {
		log.Printf("find user by name %q error: %v", register.UserName, err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "server error"})
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
	log.Printf("login params:%+v\n", login)
	user, err := dao.UserLogin(login.UserName, login.Password)
	log.Printf("search user: %+v\n", user)
	if user.Id == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "user not found"})
		return
	}
	if err != nil {
		log.Println("login Error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	resp := LoginResponse{
		Id:       user.Id,
		UserName: user.UserName,
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "user": resp})
}
