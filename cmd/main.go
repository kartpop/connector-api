package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sede-x/gopoc-connector/pkg/controllers"
	"github.com/sede-x/gopoc-connector/pkg/controllers/graphqlapi"
	"github.com/sede-x/gopoc-connector/pkg/controllers/restapi"
	"github.com/sede-x/gopoc-connector/pkg/data/kafkaq"
	"github.com/sede-x/gopoc-connector/pkg/data/postgres"
	"github.com/sede-x/gopoc-connector/pkg/logic"
)

func main() {
	// setup DB
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_CONTAINER_NAME"),
		os.Getenv("DB_CONTAINER_PORT"),
		os.Getenv("DB_NAME"),
	)
	pgdb, err := postgres.New(dbURL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// setup Kafka
	kafka, err := kafkaq.New()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// setup logic
	conlogic := &logic.Connector{DB: pgdb, MQ: kafka}

	// setup server based on APITYPE set in .env
	var server controllers.APIServer
	switch apiType := os.Getenv("APITYPE"); apiType {
	case "restapi":
		server = &restapi.Server{ConnectorLogic: conlogic}
	case "graphqlapi":
		server = &graphqlapi.Server{ConnectorLogic: conlogic}
	}

	// start server
	serverURL := fmt.Sprintf(":%s", os.Getenv("SERVER_HOST_PORT"))
	server.Start(serverURL)
}
