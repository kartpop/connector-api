package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sede-x/gopoc-connector/pkg/controllers"
	"github.com/sede-x/gopoc-connector/pkg/controllers/graphqlapi"
	"github.com/sede-x/gopoc-connector/pkg/controllers/restapi"
	"github.com/sede-x/gopoc-connector/pkg/data/postgres"
	"github.com/sede-x/gopoc-connector/pkg/logic"
)

func main() {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err.Error())
	}

	// load DB
	dbURL := os.Getenv("DBURL")
	pgdb, err := postgres.New(dbURL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// setup logic
	conlogic := &logic.Connector{DB: pgdb}

	// setup server based on set environment
	var server controllers.APIServer
	switch apiType := os.Getenv("APITYPE"); apiType {
	case "restapi":
		server = &restapi.Server{ConnectorLogic: conlogic}
	case "graphqlapi":
		server = &graphqlapi.Server{ConnectorLogic: conlogic}
	}

	// start server
	serverURL := os.Getenv("SERVERURL")
	server.Start(serverURL)
}
