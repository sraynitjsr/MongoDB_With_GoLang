package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type MongoDBClient struct {
	client *mongo.Client
}

func NewMongoDBClient() (*MongoDBClient, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	return &MongoDBClient{client}, nil
}

func (c *MongoDBClient) CreateUser(user *User) error {
	collection := c.client.Database("mydb").Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (c *MongoDBClient) GetUserByEmail(email string) (*User, error) {
	var user User
	collection := c.client.Database("mydb").Collection("users")
	err := collection.FindOne(context.Background(), map[string]string{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func main() {
	client, err := NewMongoDBClient()
	if err != nil {
		log.Fatal(err)
	}

	user := &User{
		Name:     "Subhradeep Ray",
		Email:    "subhradeepray2017@gmail.com",
		Password: "won't tell you",
	}
	err = client.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	foundUser, err := client.GetUserByEmail("subhradeepray2017@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found user: %+v\n", foundUser)
}
