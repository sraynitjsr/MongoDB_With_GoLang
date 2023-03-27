package main

import (
	"fmt"
	"log"
	"net/http"

	controllers "github.com/sraynitjsr/controllers"
)

func main() {
	fmt.Println("Starting Server")
	mdb := controllers.NewController()
	mdb.Router.GET("/", mdb.Index)
	mdb.Router.GET("/users", mdb.GetUsers)
	mdb.Router.GET("/users/{userId}", mdb.GetUserById)
	mdb.Router.POST("/users/{userId}", mdb.CreateUser)
	mdb.Router.DELETE("/users/{userId}", mdb.DeleteUser)
	log.Fatal(http.ListenAndServe(":8080", mdb.Router))
}
