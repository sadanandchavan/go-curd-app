package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sadanandchavan/go-crud-app/config"
	"github.com/sadanandchavan/go-crud-app/models"
)

// Create Property
/*
func CreateProperty(c *gin.Context) {
	var property models.Property
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&property)
	c.JSON(http.StatusOK, property)
}
*/
func CreateProperty(c *gin.Context) {
	var property models.Property

	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&property).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, property)
}

// Get All Properties
func GetProperties(c *gin.Context) {
	var properties []models.Property
	config.DB.Find(&properties)
	c.JSON(http.StatusOK, properties)
}

// Get Single Property
func GetProperty(c *gin.Context) {
	id := c.Param("id")
	var property models.Property
	if err := config.DB.First(&property, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}
	c.JSON(http.StatusOK, property)
}

// Update Property
func UpdateProperty(c *gin.Context) {
	id := c.Param("id")
	var property models.Property
	if err := config.DB.First(&property, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&property)
	c.JSON(http.StatusOK, property)
}

// Delete Property
func DeleteProperty(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Property{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
