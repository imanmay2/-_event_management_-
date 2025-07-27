package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *gin.Engine, db *mongo.Client) {
	router.POST("/register", func(c *gin.Context) {
		controllers.Register(c, db)
	})
	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})

	router.POST("/events", func(c *gin.Context) {
		controllers.CreateEvent(c, db)
	})
	router.POST("/events/:id/attend", func(c *gin.Context) {
		controllers.AttendEvent(c, db)
	})
}
