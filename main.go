package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	controllers "github.com/sraynitjsr/controllers"
)

func main() {
	fmt.Println("Starting Server")
	router := httprouter.New()
	router.GET("/", controllers.Index)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/{userId}", controllers.GetUserById)
	router.POST("/users/{userId}", controllers.CreateUser)
	log.Fatal(http.ListenAndServe(":8080", router))
}
