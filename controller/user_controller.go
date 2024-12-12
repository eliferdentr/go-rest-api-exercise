package controller

import (
	"fmt"
	"net/http"

	"eliferden.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if (err != nil) {
		context.JSON(http.StatusBadRequest, gin.H{"message:" : "Could not parse the request data"})
		return
	}

	err = user.Save()
	if (err != nil) {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message:" : "There was an error while signing up"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message" : "Sign-up is successful"})
}