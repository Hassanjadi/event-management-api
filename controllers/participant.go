package controllers

import (
	"event-management-api/database"
	"event-management-api/repository"
	"event-management-api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllParticipant(c *gin.Context) {
    var (
       result gin.H
    )

    participant, err := repository.GetAllParticipant(database.DbConnection)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": participant,
       }
    }

    c.JSON(http.StatusOK, result)
}

func GetParticipant(c *gin.Context) {
	var (
		result gin.H
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		result = gin.H{
			"result": "Invalid participant ID format",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	participant, err := repository.GetParticipant(database.DbConnection, id)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": participant,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertParticipant(c *gin.Context) {
	var participant structs.EventParticipant

	err := c.BindJSON(&participant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	err = repository.InsertParticipant(database.DbConnection, &participant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, participant)
}

func UpdateParticipant(c *gin.Context) {
	var participant structs.EventParticipant
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Invalid participant ID format"})
		return
	}

	err = c.BindJSON(&participant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	participant.ID = id

	err = repository.UpdateParticipant(database.DbConnection, participant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}

	c.JSON(http.StatusOK, participant)
}

func DeleteParticipant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Invalid participant ID format"})
		return
	}

	err = repository.DeleteParticipant(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}

    c.JSON(http.StatusOK, gin.H{
        "message": "Delete successfully",
    })
}
