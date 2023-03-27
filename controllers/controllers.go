package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sraynitjsr/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client     *mongo.Client
	Collection *mongo.Collection
	Router     *httprouter.Router
}

func NewController() *MongoDB {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ := mongo.Connect(context.Background(), clientOptions)
	client.Ping(context.Background(), nil)
	collection := client.Database("mydb").Collection("mycollection")
	defer client.Disconnect(context.Background())
	return &MongoDB{client, collection, httprouter.New()}
}

func (mdb *MongoDB) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to API building in Golang Using MongoDB\n")
}

func (mdb *MongoDB) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Getting All Users\n")
	cursor, _ := mdb.Collection.Find(context.Background(), mdb.Client)
	var users []models.User
	cursor.All(context.Background(), &users)
	fmt.Fprint(w, users)
}

func (mdb *MongoDB) GetUserById(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to API building in Golang\n")
}

func (mdb *MongoDB) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Creatng one user\n")
	user := &models.User{
		Name:   "A",
		Age:    25,
		Gender: "Male",
	}
	mdb.Collection.InsertOne(context.Background(), user)
}

func (mdb *MongoDB) DeleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to API building in Golang\n")
}
