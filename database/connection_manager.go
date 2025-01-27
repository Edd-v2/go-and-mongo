package database

import (
	"context"
	"log"
	"time"

	"go-and-mongo/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseService struct {
	DBClient     *mongo.Client
	DBName       string
	Timeout      int
	DBNamePrefix string
}

var dbService *DatabaseService

func NewGlobalDBService(cfg config.DbConfig) *DatabaseService {

	clientOptions := options.Client().
		ApplyURI(cfg.DBConnectionUri).
		SetMaxConnIdleTime(10 * time.Second).
		SetMaxPoolSize(20)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	// Connect to Mongo
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the db
	err = dbClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	// Initialize and return the DatabaseService
	log.Println("Successfully connected to MongoDB")

	dbService = &DatabaseService{
		DBClient:     dbClient,
		DBName:       cfg.DBName,
		Timeout:      cfg.Timeout,
		DBNamePrefix: cfg.DBNamePrefix,
	}

	return dbService
}

func GetDB() *mongo.Database {
	if dbService == nil || dbService.DBClient == nil {
		log.Fatal("Database is not initialized. Call NewGlobalDBService first.")
	}
	return dbService.DBClient.Database(dbService.DBName)
}

func GetCollection(collectionName string) *mongo.Collection {
	return GetDB().Collection(collectionName)
}

// Safely closes db connection
func Disconnect() {
	if dbService != nil && dbService.DBClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(dbService.Timeout)*time.Second)
		defer cancel()

		if err := dbService.DBClient.Disconnect(ctx); err != nil {
			log.Fatalf("Error while disconnecting from the database: %v", err)
		}
		log.Println("Database connection closed successfully")
	}
}
