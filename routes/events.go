package routes

import (
	"gokul-mylsami/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context){
	events,err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events ", "error" : err.Error()})
		return
	}

	context.JSON(http.StatusOK,events)
}

func getEvent(context *gin.Context){
	eventId,err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event", "error" : err.Error()})
		return
	}

	event,err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error" : err.Error()})
	}

	context.JSON(http.StatusOK,event)
}

func createEvent(context *gin.Context){

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse the data"})
		return
	}

	event.UserID = context.GetInt64("userId")

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Create events"})
		return 
	}

	context.JSON(http.StatusCreated,gin.H{"message": "Event created!", "event": event})
}

func updateEventById(context *gin.Context){
	eventId,err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event", "error" : err.Error()})
		return
	}

	event,err := models.GetEventById(eventId)
	userID := context.GetInt64("userId")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the events", "error" : err.Error()})
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized,gin.H{"message": "Not autherized to update"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Create events"})
		return 
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Update events"})
		return 
	}

	context.JSON(http.StatusOK,gin.H{"message": "Event Update Successfully"})
}

func deleteEventById(context *gin.Context){
	eventId,err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event", "error" : err.Error()})
		return
	}

	event,err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the events", "error" : err.Error()})
	}


	userID := context.GetInt64("userId")

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized,gin.H{"message": "Not autherized to Delete"})
		return
	}
	
	err = event.Delete()

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Delete events"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message": "Deleted event successfully !"})
}