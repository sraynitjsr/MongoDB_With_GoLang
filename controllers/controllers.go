package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

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
	fmt.Fprint(w, "Welcome to API building in Golang\n")
}

func (mdb *MongoDB) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to API building in Golang\n")
}

func (mdb *MongoDB) GetUserById(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to API building in Golang\n")
}

func (mdb *MongoDB) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to API building in Golang\n")
}

func (mdb *MongoDB) DeleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to API building in Golang\n")
}
