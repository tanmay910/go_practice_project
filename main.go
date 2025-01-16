package main

import (
	"fmt"
	"net/http"
	"example.com/rest-api/models"
	"example.com/rest-api/db"
	"github.com/gin-gonic/gin"
)


func main() {
	
	server := gin.Default()

	db.InitDB()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")

}

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
