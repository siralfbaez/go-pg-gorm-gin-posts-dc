package main

import (
	"github.com/siralfbaez/go-gin-gorm-docker-pg/initializers"
	"github.com/siralfbaez/go-gin-gorm-docker-pg/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnToDb()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		return
	}
}
