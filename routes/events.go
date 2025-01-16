

package routes

import (
	"strconv"
	"fmt"
	"net/http"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get events", "error": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	fmt.Println("********************************")	
	fmt.Println("sdsd %v",err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not bind JSON"})
		return
	} 
	event.ID = 1
	event.UserID = 1
	event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event. Try again"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func getEvent(context *gin.Context) {
	eventId , err:= strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event", "error": err})
		return
	}
	context.JSON(http.StatusOK, event)
}


func updateEvent(context *gin.Context) {

	eventId , err:= strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	 _, err = models.GetEventById(eventId)


	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not Fetch  event", "error": err})
		return
	} 

	var updateEvent models.Event
	err =  context.ShouldBindJSON(&updateEvent)


	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data", "error": err})
		return
	}

	updateEvent.ID = eventId
	err = updateEvent.Update()

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event", "error": err})
		return
	}


	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})


}

func deleteEvent(context *gin.Context) {


	eventId , err:= strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)


	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not Fetch  event", "error": err})
		return
	} 

	err = event.Delete()

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})


}
