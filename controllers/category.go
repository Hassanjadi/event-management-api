package controllers

import (
	"event-management-api/database"
	"event-management-api/repository"
	"event-management-api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
    var (
       result gin.H
    )

    category, err := repository.GetAllCategory(database.DbConnection)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": category,
       }
    }

    c.JSON(http.StatusOK, result)
}

func GetCategory(c *gin.Context) {
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

    category, err := repository.GetCategory(database.DbConnection, id)
    if err != nil {
        result = gin.H{
            "result": err.Error(),
        }
    } else {
        result = gin.H{
            "result": category,
        }
    }

    c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
    var category structs.Category

    err := c.BindJSON(&category)
    if err != nil {
       panic(err)
    }

    err = repository.InsertCategory(database.DbConnection, &category)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
    var category structs.Category
    id, _ := strconv.Atoi(c.Param("id"))

    err := c.BindJSON(&category)
    if err != nil {
       panic(err)
    }

    category.ID = id

    err = repository.UpdateCategory(database.DbConnection, category)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
    var category structs.Category
    id, _ := strconv.Atoi(c.Param("id"))

    category.ID = id
    err := repository.DeleteCategory(database.DbConnection, category)
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