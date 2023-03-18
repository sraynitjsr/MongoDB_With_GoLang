package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

var client *mongo.Client

func main() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(context.Background(), clientOptions)

	// Initialize routes
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/create", createUser)
	http.HandleFunc("/users/update", updateUser)
	http.HandleFunc("/users/delete", deleteUser)

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// GET all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []User
	collection := client.Database("testdb").Collection("users")
	ctx := context.Background()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// POST a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	json.NewDecoder(r.Body).Decode(&user)
	collection := client.Database("testdb").Collection("users")
	ctx := context.Background()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(result)
}

// PUT update an existing user
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	json.NewDecoder(r.Body).Decode(&user)
	collection := client.Database("testdb").Collection("users")
	ctx := context.Background()
	id, _ := primitive.ObjectIDFromHex(user.ID.Hex())
	filter := bson.M{"_id": id}

	update := bson.D{
		{"$set", bson.D{
			{"name", user.Name},
			{"email", user.Email},
			{"password", user.Password},
		}},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(result)
}

// DELETE an existing user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	id, _ := primitive.ObjectIDFromHex(params.Get("id"))
	collection := client.Database("testdb").Collection("users")
	ctx := context.Background()

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(result)
}
