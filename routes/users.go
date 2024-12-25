package routes

import (
	"gokul-mylsami/rest-api/models"
	"gokul-mylsami/rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse the data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not save user."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully"})
}

func login(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse the data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message" : "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email,user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Login successfull","token":token})
}