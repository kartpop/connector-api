package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kartpop/connector-api/pkg/controllers"
	"github.com/kartpop/connector-api/pkg/controllers/graphqlapi"
	"github.com/kartpop/connector-api/pkg/controllers/restapi"
	"github.com/kartpop/connector-api/pkg/data/kafkaq"
	"github.com/kartpop/connector-api/pkg/data/postgres"
	"github.com/kartpop/connector-api/pkg/logic"
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
