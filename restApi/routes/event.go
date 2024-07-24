package routes

import (
	"net/http"
	"strconv"

	"github.com/abcom/restApi/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	userId := context.GetInt64("userId")
	// event.ID = 1
	event.UserID = userId
	err = models.SaveEvent(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse data"})
		return
	}
	context.JSON(http.StatusOK, event)

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch data"})
		return
	}
	userId := context.GetInt64("userId")
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch data"})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to update this event"})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch data"})
		return
	}
	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this event"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete data"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
