package alts

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database string
)

const (
	// environment variables
	mongoDBConnectionStringEnvVarName = "DATABASE_MONGODB_URI"
	mongoDBDatabaseEnvVarName         = "DATABASE_MONGODB_DBNAME"
)

// connects to MongoDB
func connect() *mongo.Client {
	mongoDBConnectionString := os.Getenv(mongoDBConnectionStringEnvVarName)
	if mongoDBConnectionString == "" {
		log.Fatal("missing environment variable: ", mongoDBConnectionStringEnvVarName)
	}

	database = os.Getenv(mongoDBDatabaseEnvVarName)
	if database == "" {
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
