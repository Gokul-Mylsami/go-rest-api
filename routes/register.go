package routes

import (
	"gokul-mylsami/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId,err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event", "error" : err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Could not fetch events"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Could not register user for event"})
		return
	}

	context.JSON(http.StatusCreated,gin.H{"message": "Successfully registered user to the event"})
}

func cancelRegisteration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId,err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event", "error" : err.Error()})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegisteration(userId)


	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Could not cancel registeration" + err.Error()})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message": "Cancelled"})
}