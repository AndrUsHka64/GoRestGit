package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

// @Summary Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {array} Product
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

// @Summary Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body Product true "Product to create"
// @Success 200 {object} Product
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&product)
	c.JSON(http.StatusOK, product)
}
