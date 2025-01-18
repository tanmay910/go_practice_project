package routes


import (

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticate := server.Group("/")
	authenticate.Use(middleware.Authenticate)
	authenticate.POST("/events", createEvent)
	authenticate.PUT("/events/:id", updateEvent)
	authenticate.DELETE("/events/:id", deleteEvent)
	authenticate.POST("/events/:id/register", registerForEvent)
	authenticate.DELETE("/events/:id/unregister", unregisterForEvent)

	// server.POST("/events", middleware.Authenticate,createEvent)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login",login)


	

}
