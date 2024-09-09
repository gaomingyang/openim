package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头中的Bearer Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Bearer Token"})
			c.Abort()
			return
		}

		// 验证Token
		// 移除Bearer前缀，保留令牌部分
		// token = strings.Replace(authHeader, "Bearer ", "", 1)
		tokenString := authHeader[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt.secretKey")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Bearer Token"})
			c.Abort()
			return
		}

		// Token验证通过，将用户信息存储在上下文中
		c.Set("user", token)

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Bearer Token"})
			c.Abort()
			return
		}

		// fmt.Printf("%+v\n", claims)
		userID := claims["user_id"]

		// Token验证通过，将用户信息或其他信息存储在上下文中以备后续使用
		c.Set("user_id", userID) // 这里假设你有一个用户ID

		c.Next()
	}
}
