package controllers

import (
	"net/http"
	"pbi-task/database"
	"pbi-task/models"

	"github.com/gin-gonic/gin"
)

func postPhotos(context *gin.Context) {
	userPhoto, _ := context.Get("user")
	currentUser := userPhoto.(models.User)

	database.DB.Preload("Photos").First(&currentUser, currentUser.ID)

	if len(currentUser.Photos) >= 1 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "anda telah mengupdate foto"})
		context.Abort()
		return
	}
}
