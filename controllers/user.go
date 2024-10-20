package controllers

import (
	"event-management-api/database"
	"event-management-api/repository"
	"event-management-api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
    var (
       result gin.H
    )

    user, err := repository.GetAllUser(database.DbConnection)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": user,
       }
    }

    c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
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

    user, err := repository.GetUser(database.DbConnection, id)
    if err != nil {
        result = gin.H{
            "result": err.Error(),
        }
    } else {
        result = gin.H{
            "result": user,
        }
    }

    c.JSON(http.StatusOK, result)
}

func InsertUser(c *gin.Context) {
    var user structs.User

    err := c.BindJSON(&user)
    if err != nil {
       panic(err)
    }

    err = repository.InsertUser(database.DbConnection, &user)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
    var user structs.User
    id, _ := strconv.Atoi(c.Param("id"))

    err := c.BindJSON(&user)
    if err != nil {
       panic(err)
    }

    user.ID = id

    err = repository.UpdateUser(database.DbConnection, user)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
    var user structs.User
    id, _ := strconv.Atoi(c.Param("id"))

    user.ID = id
    err := repository.DeleteUser(database.DbConnection, user)
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
