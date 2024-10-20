package controllers

import (
	"event-management-api/database"
	"event-management-api/repository"
	"event-management-api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllReview(c *gin.Context) {
    var (
       result gin.H
    )

    review, err := repository.GetAllReview(database.DbConnection)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": review,
       }
    }

    c.JSON(http.StatusOK, result)
}

func GetReview(c *gin.Context) {
	var (
		result gin.H
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		result = gin.H{
			"result": "Invalid review ID format",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	review, err := repository.GetReview(database.DbConnection, id)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": review,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertReview(c *gin.Context) {
	var review structs.Review

	err := c.BindJSON(&review)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	err = repository.InsertReview(database.DbConnection, &review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}

func UpdateReview(c *gin.Context) {
	var review structs.Review
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Invalid review ID format"})
		return
	}

	err = c.BindJSON(&review)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	review.ID = id

	err = repository.UpdateReview(database.DbConnection, review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

func DeleteReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Invalid review ID format"})
		return
	}

	err = repository.DeleteReview(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}

    c.JSON(http.StatusOK, gin.H{
        "message": "Delete successfully",
    })
}
