package main

import (
	"context"
	"go-and-mongo/config"
	"go-and-mongo/database"
	"go-and-mongo/internal/service"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	config.LoadConfiguration()

	dbService := database.NewGlobalDBService(config.ServerConf.DbConfig)
	if dbService == nil || dbService.DBClient == nil {
		log.Fatal("Database connection failed.")
	}

	// user service
	userService := service.NewUserService(database.GetDB())

	// TEST
	users := []bson.M{
		{"name": "John Doe", "age": 30, "email": "johndoe@example.com", "job": "Software Engineer"},
		{"name": "Jane Smith", "age": 25, "email": "janesmith@example.com", "job": "Data Scientist"},
		{"name": "Mary Doe", "age": 22, "email": "marydoe@example.com", "job": "Designer"},
	}
	for _, user := range users {
		insertResult, err := userService.Create(context.Background(), user)
		if err != nil {
			log.Fatalf("Error inserting user: %v", err)
		}
		log.Printf("Inserted user with ID: %v", insertResult.InsertedID)
	}

	filter := bson.M{"name": "John Doe"}
	user, err := userService.FindOne(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error finding user: %v", err)
	}
	log.Printf("Found user: %v", user)

	update1 := bson.M{"$set": bson.M{"age": 75}}
	updateResult1, err := userService.UpdateOne(context.Background(), bson.M{"name": "John Doe"}, update1)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	log.Printf("Updated John Doe's age. Matched %v document(s) and modified %v document(s)", updateResult1.MatchedCount, updateResult1.ModifiedCount)

	update2 := bson.M{"$set": bson.M{"job": "Senior Data Scientist"}}
	updateResult2, err := userService.UpdateOne(context.Background(), bson.M{"name": "Jane Smith"}, update2)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	log.Printf("Updated Jane Smith's job. Matched %v document(s) and modified %v document(s)", updateResult2.MatchedCount, updateResult2.ModifiedCount)

	deleteFilter := bson.M{"name": "Mary Doe"}
	deleteResult, err := userService.DeleteOne(context.Background(), deleteFilter)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
	log.Printf("Deleted %v document(s)", deleteResult.DeletedCount)

	defer database.Disconnect()
}
