package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string, timeout int) (*mongo.Client, context.Context, context.CancelFunc, error) {

	// Set Client Timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	// Set Client Options
	credentials := options.Credential{
		Username:      "sray",
		Password:      "password",
		AuthMechanism: "SCRAM-SHA-1",
	}
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credentials)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, cancel, err
}

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func main() {
	timeout := 30
	client, ctx, cancel, err := connect("mongodb://localhost:27017", timeout)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)

	var document = bson.D{
		{Key: "Roll", Value: 10},
		{Key: "Mathematics", Value: 70},
		{Key: "Science", Value: 85},
	}

	insertOneResult, err := insertOne(client, ctx, "my-mongo-db", "students", document)

	if err != nil {
		panic(err)
	}

	fmt.Println("Output of One Document Insertion")

	fmt.Println(insertOneResult.InsertedID)

	if err != nil {
		panic(err)
	}

}

// Reference Code => https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
