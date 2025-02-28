package main

import (
	"fmt"
	"goMux/internal/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")
	port := ":9860"
	StartServer(port)

	// fmt.Println("Server Started - ", port)
}

func StartServer(port string) {
	r := mux.NewRouter()
	api.RegisterRoutes(r)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(port, r))

	er := http.ListenAndServe(port, nil)
	if er != nil {
		log.Println("Error - ", er)
		// return er
	}
	// return nil
}
