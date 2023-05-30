package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siralfbaez/go-gin-gorm-docker-pg/controllers"
	"github.com/siralfbaez/go-gin-gorm-docker-pg/initializers"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnToDb()
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	err := r.Run()
	if err != nil {
		return
	}
}
