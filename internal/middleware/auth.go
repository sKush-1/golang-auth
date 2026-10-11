package middleware

import (
	"go/token"
	"golang-auth/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ctxUserIDKey = "auth.userId"
	ctxRoleKey = "auth.role"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":"Missing Authorization"

			})
			return 
		}

		parts := strings.SplitN(authHeader," ",2)
		if len(parts) != 2{
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error":"Invalid authorization format"
			})
			return
		}

		scheme := strings.TrimSpace(parts[0])
		tokenString := strings.TrimSpace(parts[1])

		if !strings.EqualFold(scheme,"Bearer"){
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error":"Authorization scheme must be bearer"
			})
			return
		}

		if tokenString == ""{
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error":"Missing token"
			})
			return
		}

		claims,err := auth.ParseToken(jwtSecret)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token"
			})
		}

		c.Set(ctxUserIDKey,claims.Subject)
		c.Set(ctxRoleKey, claims.Role)

		c.Next()


	}
}


func GetUserID(c *gin.Context) (string,bool){
	res,ok := c.Get(ctxUserIDKey)
	if !ok {
		return "",false
	}

	userID,ok := res.(string)
	return userID,ok
}

func GetRole(c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxRolekey)
	if !ok {
		return "", false
	}

	role, ok := res.(string)
	return role, ok
}