package controllers

import (
	"net/http"
	"pbi-task/auth"
	"pbi-task/database"
	"pbi-task/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindBodyWithJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var existingUser models.User
	database.DB.Where("email= ?", user.Email).First(&existingUser)
	if existingUser.ID != 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "email already exist"})
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.DB.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userID": user.ID, "email": user.Email, "username": user.Username})

}

func UserLogin(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindBodyWithJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "can't read user input"})
		context.Abort()
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed", "message": err.Error()})
		context.Abort()
		return
	}

	var existingUser models.User

	if err := database.DB.Find(&existingUser, "email= ?", user.Email).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "user not found"})
		context.Abort()
		return
	}

	errHash := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))

	if errHash != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	tokenString, _ := auth.GenerateJWT(user.ID)

	context.JSON(http.StatusAccepted, gin.H{"status": "success", "message": "berhasil login", "userID": existingUser.ID, "token": tokenString})

}

func GetUser(context *gin.Context) {
	var user []models.User
	database.DB.Preload("Photos").Find(&user)

	context.JSON(http.StatusOK, gin.H{"user data": user})
}

func GetUserById(context *gin.Context) {
	id := context.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "user not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"messaga": "found user", "user": user})
}

func UserUpdate(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindBodyWithJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "can't read user input"})
		context.Abort()
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed", "message": err.Error()})
		context.Abort()
		return
	}

	var existingUser models.User
	userID := context.Param("id")
	searchUser := database.DB.First(&existingUser, "id= ?", userID)

	if searchUser.RowsAffected == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if searchUser.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
		return
	}

	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Password = string(hash)

	if err := database.DB.Model(&existingUser).Where("id = ?", userID).Updates(&existingUser).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "berhasil", "message": "data telah diupdate", "user": existingUser})

}

func UserDelete(context *gin.Context) {
	id := context.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "failed", "message": err.Error()})
		context.Abort()
		return
	}

	if err := database.DB.Unscoped().Delete(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed", "message": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "berhasil", "message": "user telah dihapus"})
}
