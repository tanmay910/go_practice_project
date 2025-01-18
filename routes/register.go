package routes

import (
	"strconv"
	"net/http"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId , err:= strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event , err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event"})
		return
	}	
	 err = event.Register(userId)


	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "registered for event successfully"})

}

func unregisterForEvent(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId , err:= strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event , err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event"})
		return
	}
	
	
	err = event.Unregister(userId)

	 if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Registration  cancellation request failed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Unregistered for event successfully"})


	}