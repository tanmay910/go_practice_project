

package routes

import (
	"fmt"
	"net/http"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"example.com/rest-api/utlities"
)

func signup(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)

	

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not bind JSON"})
		return
	} 
	
	 err = user.Save()
	if err != nil {
		fmt.Println("Error : ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user. Try again"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})

}

func login(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"message": "could not bind JSON"})
		return
	} 
	
	err = user.ValidateCredentials()

	if err!= nil {
		fmt.Println("error: " , err)
		context.JSON(http.StatusBadRequest , gin.H{"message":"Could not authenticate user."})
		return
	}

	token , err := utlities.GenerateToken(user.Email, user.ID)
	if err != nil {
		fmt.Println("error: " , err)
		context.JSON(http.StatusInternalServerError , gin.H{"message":"Could not generate token."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Login successful!", "token": token})

}