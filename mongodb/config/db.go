package config

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const connectionString = "mongodb://127.0.0.1:27017/golang"
const dbName = "golang"
const colName = "users"

var UserCollection *mongo.Collection

// connect with mongoDB

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connectin Success")

	UserCollection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")

}
