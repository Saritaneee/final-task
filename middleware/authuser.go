package middleware

import (
	"net/http"
	"pbi-task/database"
	"pbi-task/models"

	"github.com/gin-gonic/gin"
)

func AuthUser(context *gin.Context) {
	userID := context.Param("id")
	reqID := context.GetUint("reqID")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}
	ID := user.ID

	if reqID != ID {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": "Failed", "Message": "You dont have access"})
		return
	}

	context.Next()
}
