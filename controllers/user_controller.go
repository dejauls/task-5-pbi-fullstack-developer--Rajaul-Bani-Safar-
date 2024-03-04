package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/app"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/helpers"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/database"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)


func RegisterUser(c *gin.Context) {
	var registrationRequest app.CreateUserRequest

	if err := c.ShouldBindJSON(&registrationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := app.ValidateCreateUserRequest(registrationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	err := database.DB.Where("email = ?", registrationRequest.Email).First(&existingUser).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registrationRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}

	newUser := models.User{
		Username: registrationRequest.Username,
		Email:    registrationRequest.Email,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"username": newUser.Username,
			"email":    newUser.Email,
			"password": "********",
		},
		"message": "User registered successfully",
	})
}


func LoginUser(c *gin.Context) {

	var loginRequest app.LoginUserRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := database.DB.Where("email = ?", loginRequest.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := helpers.GenerateJWTToken(user.Email, user.Username, strconv.Itoa(int(user.ID)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"username": user.Username,
			"email":    user.Email,
			"authorization": gin.H{
				"bearer": token,
			},
		},
		"message": "Login successful",
	})
}


func UpdateUser(c *gin.Context) {
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userIDFromToken, err := helpers.ExtractUserIDFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if userID != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	var updateRequest app.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	var existingUser models.User
	err = database.DB.First(&existingUser, userID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if existingUser.Email != updateRequest.Email {
		var userWithEmailExists models.User
		err = database.DB.Where("email = ?", updateRequest.Email).First(&userWithEmailExists).Error
		if err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
			return
		}
	}

	existingUser.Username = updateRequest.Username
	existingUser.Email = updateRequest.Email

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}
	existingUser.Password = string(hashedPassword)

	if err := database.DB.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "userId": userID})
}

func DeleteUser(c *gin.Context) {
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userIDFromToken, err := helpers.ExtractUserIDFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if userID != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	var user models.User
	err = database.DB.First(&user, userID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "userId": userID})
}
