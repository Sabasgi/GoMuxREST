package api

import (
	"goMux/internal/api/handlers"

	"github.com/gorilla/mux"
)

/*
	/o – Open Routes

	These are public or unauthenticated endpoints.
	Example: /o/login (Login endpoint does not require authentication).
	/c – Create/Closed Routes

	This typically represents protected endpoints that require authentication and are often used for creating resources.
	Example: /c/add/product (Only authenticated users can add products).
	/r – Read/Restricted Routes

	This usually indicates protected endpoints used for retrieving data.
	Example: /r/get/product (Authenticated users can fetch product details).
*/

func RegisterRoutes(router *mux.Router) {
	// open routes
	openRoutes := router.PathPrefix("/o").Subrouter()
	//c roued
	cRoutes := router.PathPrefix("/c").Subrouter()
	cRoutes.Use()
	//R routes - private routes
	rRoutes := router.PathPrefix("/r/c").Subrouter()
	rRoutes.Use()

	//Routes
	openRoutes.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
}
