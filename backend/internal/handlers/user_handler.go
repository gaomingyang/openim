package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"openim/internal/common"
	"openim/internal/common/logger"
	"openim/internal/dao"
	"time"
)

type LoginRequest struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginResponse struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

var log *logrus.Logger

func init() {
	log = logger.Log
}

// login
func LoginHandler(c *gin.Context) {
	var login LoginRequest
	if err := c.ShouldBind(&login); err != nil {
		common.BadRequest(c, "parameters error")
		// c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "parameters error"})
		return
	}
	log.Info("login receive:%+v", login)
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

	token, err := CreateToken(user.Id)
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
	// tokenString := c.GetHeader("Authorization")
	// if tokenString == "" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
	// 	return
	// }

	// 移除Bearer前缀，保留令牌部分
	// tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(viper.GetString("jwt.secretKey")), nil
	// })

	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
	// 	return
	// }

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	fmt.Printf("%+v\n", claims)
	// 	userId := claims["user_id"]
	// 	c.JSON(http.StatusOK, gin.H{"message": "ok", "user_id": userId})
	// } else {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
	// }

	userID, _ := c.Get("user_id")

	c.JSON(http.StatusOK, gin.H{"message": "ok", "the_user_id": userID})
}

func CreateToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	secretKey := viper.GetString("jwt.secretKey")
	return token.SignedString([]byte(secretKey))
}
