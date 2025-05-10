package controllers

import (
	"fmt"
	"net/http"

	models "github.com/VladVozhzhov/inventory-managment-api/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Add a new product to the inventory
func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save product to the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added successfully", "product": product})
}

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// Update the stock for a specific product (quantity)
func UpdateProduct(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Product ID is required"})
		return
	}

	var product models.Product
	if err := DB.First(&product, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var updateData models.Product
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if provided
	if updateData.Name != "" {
		product.Name = updateData.Name
	}
	if updateData.SKU != "" {
		product.SKU = updateData.SKU
	}
	if updateData.Category != "" {
		product.Category = updateData.Category
	}
	if updateData.Supplier != "" {
		product.Supplier = updateData.Supplier
	}
	if updateData.Description != "" {
		product.Description = updateData.Description
	}
	// If quantity is set, update and log stock change
	quantityChange := updateData.Quantity - product.Quantity
	if updateData.Quantity != 0 && quantityChange != 0 {
		product.Quantity = updateData.Quantity
		if err := DB.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update product"})
			return
		}
		logStockTransaction(c, product.ID, quantityChange, "stock update")
	} else {
		if err := DB.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update product"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

// Logs a stock transaction (e.g., adding/removing stock)
func logStockTransaction(c *gin.Context, productID string, change int, reason string) {
	user, _ := c.Get("user")
	claims := user.(jwt.MapClaims)
	userID := claims["sub"].(string)

	transaction := models.Stock{
		ProductID: productID,
		Change:    change,
		Reason:    reason,
		UserID:    userID,
	}

	// Log the transaction to the database
	if err := DB.Create(&transaction).Error; err != nil {
		fmt.Println("Error logging stock transaction:", err)
	}
}
