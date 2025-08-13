/*
* Copyright 2019 Celona, Inc. or its affiliates. All Rights Reserved.
 */

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "net/http/pprof"

	"github.com/go-openapi/loads"
	"github.com/munishchaudhary/book-store-app/src/db"
	"github.com/munishchaudhary/book-store-app/src/generated/restapi"
	"github.com/munishchaudhary/book-store-app/src/generated/restapi/operations"
	"github.com/munishchaudhary/book-store-app/src/handlers"
)

func main() {

	// init the server
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Printf("Error: %v loading swagger", err)
		os.Exit(2)
	}

	api := operations.NewBookStoreAPI(swaggerSpec)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	// Initialize database connection
	db.Connect()

	// Initialize book handler with database
	bookDB := db.GetBookDB()
	bookHandler := handlers.NewBookHandler(bookDB)
	handlers.SetGlobalBookHandler(bookHandler)

	// Register API handlers
	handlers.RegisterAPIHandlers(api)
	server.ConfigureAPI()

	// start the server now
	setServer(server, api)
	if err = server.Serve(); err != nil {
		log.Printf("Error %v while serving requests", err)
		os.Exit(3)
	}
}

func setServer(server *restapi.Server, api *operations.BookStoreAPI) {
	schemes := os.Getenv("SVC_SCHEMES")
	server.Host = os.Getenv("SVC_HOSTNAME")
	server.EnabledListeners = strings.Split(schemes, ",")
	portStr := os.Getenv("SVC_PORT")
	if portStr == "" {
		log.Println("SVC_PORT environment variable is not set. Using default port 8080.")
		portStr = "8080"
	}
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("Error converting port to integer: %v\n", err)
		return
	}
	server.Port = portInt
}
