package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"product-management/pkg/controllers"
	"product-management/pkg/repositories"
)

func Auth(mariadb *gorm.DB, engine *gin.Engine) {

	// Homework = Check what the Engin
	userRepository, _ := repositories.NewUserRepository(mariadb)
	authController := controllers.NewAuthenticateController(userRepository)

	product := engine.Group("/auth")

	product.POST("/", authController.Authenticate())
}
