package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DbName string
)

const (
	// environment variables
	mongoDBConnectionStringEnvVarName = "DATABASE_MONGODB_URI"
	mongoDBDatabaseEnvVarName         = "DATABASE_MONGODB_DBNAME"
)

// connects to MongoDB
func Connect() *mongo.Client {
	mongoDBConnectionString := os.Getenv(mongoDBConnectionStringEnvVarName)
	if mongoDBConnectionString == "" {
		log.Fatal("missing environment variable: ", mongoDBConnectionStringEnvVarName)
	}

	DbName = os.Getenv(mongoDBDatabaseEnvVarName)
	if DbName == "" {
		log.Fatal("missing environment variable: ", mongoDBDatabaseEnvVarName)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoDBConnectionString).SetDirect(true)
	c, err := mongo.NewClient(clientOptions)

	err = c.Connect(ctx)

	if err != nil {
		log.Fatalf("failed to initialize connection %v", err)
	}
	err = c.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("failed to connect %v", err)
	}
	return c
}
