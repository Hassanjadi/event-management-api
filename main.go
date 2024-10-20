package main

import (
	"database/sql"
	"event-management-api/controllers"
	"event-management-api/database"
	"event-management-api/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var (
    DB  *sql.DB
    err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

    psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
    os.Getenv("DB_HOST"),
    os.Getenv("DB_PORT"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
    )

    DB, err = sql.Open("postgres", psqlInfo)
    defer DB.Close()
    err = DB.Ping()
    if err != nil {
       panic(err)
    }

    database.DBMigrate(DB)

    router := gin.Default()
    router.Use(middleware.Auth())

    router.GET("/api/users", controllers.GetAllUser)
    router.POST("/api/users", controllers.InsertUser)
    router.GET("/api/users/:id", controllers.GetUser)
    router.PUT("/api/users/:id", controllers.UpdateUser)
    router.DELETE("/api/users/:id", controllers.DeleteUser)
    
    router.GET("/api/categories", controllers.GetAllCategory)
    router.POST("/api/categories", controllers.InsertCategory)
    router.GET("/api/categories/:id", controllers.GetCategory)
    router.PUT("/api/categories/:id", controllers.UpdateCategory)
    router.DELETE("/api/categories/:id", controllers.DeleteCategory)

    router.GET("/api/events", controllers.GetAllEvent)
    router.POST("/api/events", controllers.InsertEvent)
    router.GET("/api/events/:id", controllers.GetEvent)
    router.PUT("/api/events/:id", controllers.UpdateEvent)
    router.DELETE("/api/events/:id", controllers.DeleteEvent)

    router.GET("/api/reviews", controllers.GetAllReview)
    router.POST("/api/reviews", controllers.InsertReview)
    router.GET("/api/reviews/:id", controllers.GetReview)
    router.PUT("/api/reviews/:id", controllers.UpdateReview)
    router.DELETE("/api/reviews/:id", controllers.DeleteReview)

    router.GET("/api/participants", controllers.GetAllParticipant)
    router.POST("/api/participants", controllers.InsertParticipant)
    router.GET("/api/participants/:id", controllers.GetParticipant)
    router.PUT("/api/participants/:id", controllers.UpdateParticipant)
    router.DELETE("/api/participants/:id", controllers.DeleteParticipant)

    router.Run(":8080")
}