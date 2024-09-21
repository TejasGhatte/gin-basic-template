package main

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"gin-app/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.DBMigrate()
	initializers.AddLogger()
	initializers.ConnectToCache()
}

func main() {
	app := gin.Default()

	app.GET("/",Hola)

	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Run(":" + initializers.CONFIG.PORT)
}

func Hola(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hola Amigo",
	})
}