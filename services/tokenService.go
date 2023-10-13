package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func refreshTokenHandler(c *gin.Context) {
	oldTokenString := c.GetHeader("Authorization")

	if oldTokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "没有提供令牌"})
		return
	}

	oldToken, err := jwt.Parse(oldTokenString, func(token *jwt.Token) (interface{}, error) {
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
			if int64(exp)-time.Now().Unix() > 60 { // 60秒内不自动刷新
				c.JSON(http.StatusOK, gin.H{"message": "令牌尚未过期"})
				return
			}
		}
	}

	// 如果旧令牌需要刷新，生成一个新令牌
	if userId, ok := claims["user_id"].(int64); ok {
		newToken, err := createToken(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建新令牌"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": newToken})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "无法刷新令牌"})
}
