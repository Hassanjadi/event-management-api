package controllers

import (
	"event-management-api/database"
	"event-management-api/repository"
	"event-management-api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllEvent(c *gin.Context) {
    var (
       result gin.H
    )

    event, err := repository.GetAllEvent(database.DbConnection)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": event,
       }
    }

    c.JSON(http.StatusOK, result)
}

func GetEvent(c *gin.Context) {
    var (
        result gin.H
    )

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        result = gin.H{
            "result": "Invalid ID format",
        }
        c.JSON(http.StatusBadRequest, result)
        return
    }

    event, err := repository.GetEvent(database.DbConnection, id)
    if err != nil {
        result = gin.H{
            "result": err.Error(),
        }
    } else {
        result = gin.H{
            "result": event,
        }
    }

    c.JSON(http.StatusOK, result)
}

func InsertEvent(c *gin.Context) {
    var event structs.Event

    err := c.BindJSON(&event)
    if err != nil {
       panic(err)
    }

    err = repository.InsertEvent(database.DbConnection, &event)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, event)
}

func UpdateEvent(c *gin.Context) {
    var event structs.Event
    id, _ := strconv.Atoi(c.Param("id"))

    err := c.BindJSON(&event)
    if err != nil {
       panic(err)
    }

    event.ID = id

    err = repository.UpdateEvent(database.DbConnection, event)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, event)
}

func DeleteEvent(c *gin.Context) {
    var event structs.Event
    id, _ := strconv.Atoi(c.Param("id"))

    event.ID = id
    err := repository.DeleteEvent(database.DbConnection, event)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Delete successfully",
    })
}