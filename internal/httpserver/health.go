package httpserver

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

func health(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"service": "go-auth",
		"timestamps": time.Now().Unix(),
	})

}