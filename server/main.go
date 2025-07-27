package main

import (
	"context"
	"log"
	"time"

	"server/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func initDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("✅ MongoDB connected")
	return client
}

func main() {
	DB = initDB()
	app := gin.Default()
	routes.SetupRoutes(app, DB)
	app.Run(":8080")
}
