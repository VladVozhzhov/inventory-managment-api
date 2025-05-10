package routes

import (
	"github.com/VladVozhzhov/inventory-managment-api/controllers"
	middlewares "github.com/VladVozhzhov/inventory-managment-api/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	controllers.DB = db

	// Public routes (no JWT required)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	router.GET("/products", controllers.GetAllProducts)

	// Protected routes (JWT required)
	authorized := router.Group("/admin")
	authorized.Use(middlewares.JWTVerify())
	authorized.POST("/products", controllers.AddProduct)
	authorized.PUT("/products", controllers.UpdateProduct)
}
