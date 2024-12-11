package controller

import (
	"eliferden.com/restapi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not fetch events, try again later."})
		log.Fatal(err)
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not parse event id."})
		log.Fatal(err)
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "There was an error while fetching the event"})
		log.Fatal(err)
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
	event.UserID = 1

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
		log.Fatal(err)
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "There was an error while fetching the event"})
		log.Fatal(err)
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
		log.Fatal(err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message: ": "Event updated successfully"})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not parse event id."})
		log.Fatal(err)
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "There was an error while fetching the event"})
		log.Fatal(err)
		return
	}
	err = event.DeleteEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message: ": "There was an error while deleting the event"})
		log.Fatal(err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message: ": "Event deleted successfully"})

}
