package main

import (
	"github.com/VladVozhzhov/inventory-managment-api/config"
	models "github.com/VladVozhzhov/inventory-managment-api/model"
	"github.com/VladVozhzhov/inventory-managment-api/routes"

	// "github.com/VladVozhzhov/inventory-managment-api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDatabase()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Stock{})

	// utils.Token()

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	routes.SetupRoutes(router, db)
	router.Run(":3500")
}
