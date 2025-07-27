package controllers

import (
	"context"
	"server/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getClaims(c *gin.Context) (string, string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", "", fmt.Errorf("missing token")
	}
	tokenStr := strings.Split(authHeader, "Bearer ")[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		role := claims["role"].(string)
		return email, role, nil
	}
	return "", "", err
}

func CreateEvent(c *gin.Context, db *mongo.Client) {
	email, role, err := getClaims(c)
	if err != nil || role != "organizer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized or invalid token"})
		return
	}

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.OrganizerID = email
	event.Attendees = []string{}

	_, err = db.Database("eventdb").Collection("events").InsertOne(context.TODO(), event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully"})
}

func AttendEvent(c *gin.Context, db *mongo.Client) {
	email, role, err := getClaims(c)
	if err != nil || role != "attendee" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized or invalid token"})
		return
	}
	eventID := c.Param("id")

	filter := bson.M{"_id": eventID}
	update := bson.M{"$addToSet": bson.M{"attendees": email}}

	_, err = db.Database("eventdb").Collection("events").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to attend event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully joined event"})
}
