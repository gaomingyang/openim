package services

import (
	"net/http"
	"openim/internal/handlers"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// 刷新token。刷新后老的在有效期内也能继续用。
func RefreshTokenHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "没有提供令牌"})
		return
	}

	// 移除Bearer前缀，保留令牌部分
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	oldToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secretKey")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "令牌无效"})
		return
	}

	// 检查是否令牌快过期
	claims, ok := oldToken.Claims.(jwt.MapClaims)
	if ok && oldToken.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp)-time.Now().Unix() > 600 { // 临过期600s内才能刷新，否则不刷新
				c.JSON(http.StatusOK, gin.H{"message": "令牌尚未过期"})
				return
			}
		}
	}

	// 如果旧令牌需要刷新，生成一个新令牌
	if userIdfloat, ok := claims["user_id"].(float64); ok {
		// fmt.Printf("userId:%+v\n", userIdfloat)
		userId := int64(userIdfloat)
		newToken, err := handlers.CreateToken(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建新令牌"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": newToken})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "无法刷新令牌"})
}
