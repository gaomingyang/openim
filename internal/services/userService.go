package services

import (
	"fmt"
	"log"
	"net/http"
	"openim/internal/common"
	"openim/internal/dao"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	// Email    string `form:"email" json:"email" binding:"required"` //后面加上
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
	Token    string `json:"token"`
}

func createToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	secretKey := viper.GetString("jwt.secretKey")
	return token.SignedString([]byte(secretKey))
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

// login
func LoginHandler(c *gin.Context) {
	var login LoginRequest
	if err := c.ShouldBind(&login); err != nil {
		common.BadRequest(c, "parameters error")
		// c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "parameters error"})
		return
	}
	// log.Printf("login params:%+v\n", login)
	user, err := dao.UserLogin(login.UserName, login.Password)
	// log.Printf("search user: %+v\n", user)
	if user.Id == 0 {
		common.BadRequest(c, "user not found")
		// c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "user not found"})
		return
	}
	if err != nil {
		log.Println("login Error", err.Error())
		common.BadRequest(c, err.Error())
		// c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	resp := LoginResponse{
		Id:       user.Id,
		UserName: user.UserName,
	}

	token, err := createToken(user.Id)
	if err != nil {
		log.Println(err.Error())
		common.InternalServerError(c, "internal error create user token failed")
		return
	}
	resp.Token = token
	common.OK(c, resp)
	// c.JSON(http.StatusOK, gin.H{"status": "success", "user": resp})
}

// test check and parse token
func UserInfoHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secretKey")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("%+v\n", claims)
		userId := claims["user_id"]
		c.JSON(http.StatusOK, gin.H{"message": "ok", "user_id": userId})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
	}
}
