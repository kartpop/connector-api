package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sede-x/gopoc-connector/pkg/controllers"
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

	// setup logic and controller
	conlogic := &logic.Connector{DB: pgdb}
	concontroller := &controllers.ConnectorRestAPI{ConnectorLogic: conlogic}

	// start server
	serverURL := os.Getenv("SERVERURL")
	concontroller.StartServer(serverURL)
}
