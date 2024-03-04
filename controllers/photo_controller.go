package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/app"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/helpers"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/database"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/models"
	"strconv"
)

func CreatePhoto(c *gin.Context) {
	var createPhotoRequest app.CreatePhotoRequest
	if err := c.ShouldBindJSON(&createPhotoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDFromToken, err := helpers.ExtractUserIDFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID := uint(userIDFromToken)

	newPhoto := models.Photo{
		Title:    createPhotoRequest.Title,
		Caption:  createPhotoRequest.Caption,
		PhotoURL: createPhotoRequest.PhotoURL,
		UserID:   userID,
	}

	if err := database.DB.Create(&newPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":        newPhoto.ID,
			"title":     newPhoto.Title,
			"caption":   newPhoto.Caption,
			"photoUrl":  newPhoto.PhotoURL,
			"userId":    newPhoto.UserID,
			"createdAt": newPhoto.CreatedAt,
			"updatedAt": newPhoto.UpdatedAt,
		},
		"message": "Photo created successfully",
	})
}


func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	if err := database.DB.Preload("User").Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get photos"})
		return
	}

	for i := range photos {
		photos[i].User.Password = ""
	}

	c.JSON(http.StatusOK, gin.H{"data": photos, "message": "Photos retrieved successfully"})
}



func UpdatePhoto(c *gin.Context) {
	photoIDStr := c.Param("photoId")
	photoID, err := strconv.ParseUint(photoIDStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	userIDFromToken, err := helpers.ExtractUserIDFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var existingPhoto models.Photo
	err = database.DB.First(&existingPhoto, photoID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if uint64(existingPhoto.UserID) != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}
	

	var updateRequest app.UpdatePhotoRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingPhoto.Title = updateRequest.Title
	existingPhoto.Caption = updateRequest.Caption
	existingPhoto.PhotoURL = updateRequest.PhotoURL

	if err := database.DB.Save(&existingPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully", "photoId": photoID})
}


func DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoId")

	
	userIDFromToken, err := helpers.ExtractUserIDFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	
	var existingPhoto models.Photo
	if err := database.DB.Where("id = ?", photoID).First(&existingPhoto).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}
	if existingPhoto.UserID != uint(userIDFromToken) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	
	if err := database.DB.Delete(&existingPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully", "photoId": photoID})
}
