package connection

import (
	"context"

	"fmt"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"os"
)

var client *mongo.Client

func ConnectToMongoDB() *mongo.Client {

    if err := godotenv.Load(); err != nil {
        log.Println("No .env file is specified.")
    }
    if client != nil {
        return nil
    }

    uri := os.Getenv("MONGODB_URI")
    if uri == "" {
        fmt.Println("MONGODB_URI is not set")
    }

    newClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    if err != nil {
         return nil
    }

    client = newClient
    return client
}