package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"eliferden.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not fetch events, try again later."})
		 
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not parse event id."})
		 
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "There was an error while fetching the event"})
		 
		return
	}

	context.JSON(http.StatusOK, event)
}

func CreateEvent(context *gin.Context) {
	

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "There was an error while saving the event"})
		fmt.Print(err)
		return
	}
	event.UserID = context.GetInt64("userId")

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "There was an error while saving the event to the database"})
		fmt.Print(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	event , err := models.GetEventByID(eventId)

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message" : "Not authorized to update event"})
		return 
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "There was an error while fetching the event"})
		 
		return
	}
	

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "There was an error while saving the event"})
		fmt.Print(err)
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "There was an error while updating the event"})
		 
		return
	}
	context.JSON(http.StatusOK, gin.H{"message: ": "Event updated successfully"})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not parse event id."})
		 
		return
	}

	userId := context.GetInt64("userId")
	event , err := models.GetEventByID(eventId)

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message" : "Not authorized to delete event"})
		return 
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "There was an error while fetching the event"})
		 
		return
	}
	err = event.DeleteEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "There was an error while deleting the event"})
		 
		return
	}
	context.JSON(http.StatusOK, gin.H{"message: ": "Event deleted successfully"})

}

func RegisterForEvent (context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not parse event id."})
		 
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "Could not find the event."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message" : "Event is registered"})

}

func CancelRegistration (context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not parse event id."})
		 
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "Could not find the event."})
		return
	}

	err = event.Unregister(userId, eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message" : "Event is unregistered"})
}
