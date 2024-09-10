package initialisers

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDB() (*mongo.Client, error) {
    err := godotenv.Load("../../.env")
    if err != nil {
        return nil, fmt.Errorf("error loading .env file")
    }

    dbUri := os.Getenv("DB_URI")
    if dbUri == "" {
        return nil, fmt.Errorf("DB URI is not set")
    }

    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(dbUri).SetServerAPIOptions(serverAPI)

    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
        return nil, err
    }

    if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
        return nil, err
    }

    fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
    return client, nil
}

// func main() {
// 	connectToDB()
// }
