package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yosefalemu/GoCrudChallange/models"
	"github.com/yosefalemu/GoCrudChallange/utils"
	"gorm.io/gorm"
)

// Creates a person
func CreatePerson(c *gin.Context, db *gorm.DB) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}
	if err := db.Create(&person).Error; err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Failed to create person", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": person})
}

// Get all people
func GetPerson(c *gin.Context, db *gorm.DB) {
	fmt.Println("GetPerson")
	var people []models.Person
	if err := db.Find(&people).Error; err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Failed to fetch people", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": people})
}

// Get a person by ID
func GetPersonById(c *gin.Context, db *gorm.DB) {
	var person models.Person
	id := c.Param("id")
	if err := db.First(&person, "id = ?", id).Error; err != nil {
		utils.ResponseWithError(c, http.StatusNotFound, "Person not found", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": person})
}

// Update a person by ID
func UpdatePersonById(c *gin.Context, db *gorm.DB) {
	var person models.Person
	id := c.Param("id")
	if err := db.First(&person, "id = ?", id).Error; err != nil {
		utils.ResponseWithError(c, http.StatusNotFound, "Person not found", err)
		return
	}
	if err := c.ShouldBindJSON(&person); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}
	if err := db.Save(&person).Error; err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Failed to update person", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": person})
}

// Delete a person by ID
func DeletePersonById(c *gin.Context, db *gorm.DB) {
	var person models.Person
	id := c.Param("id")
	if err := db.First(&person, "id = ?", id).Error; err != nil {
		utils.ResponseWithError(c, http.StatusNotFound, "Person not found", err)
		return
	}
	if err := db.Delete(&person).Error; err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Failed to delete person", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Person deleted successfully"})
}
