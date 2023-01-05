package graph

import (
	"log"
	"strconv"

	"github.com/sede-x/gopoc-connector/pkg/controllers/graphqlapi/graph/model"
	"github.com/sede-x/gopoc-connector/pkg/models"
)

func convertToGraphConnector(con models.Connector) model.Connector {
	var connector model.Connector
	connector.ID = strconv.Itoa(con.Id)
	connector.StationID = con.StationId
	connector.ChargeSpeed = con.ChargeSpeed
	connector.Type = con.Type
	connector.Active = con.Active

	return connector
}

func convertToDBConnector(con model.Connector) models.Connector {
	var connector models.Connector
	if con.ID != "" {
		connectorId, err := strconv.Atoi(con.ID)
		if err != nil {
			log.Fatalln(err.Error())
		}
		connector.Id = connectorId
	}
	connector.StationId = con.StationID
	connector.ChargeSpeed = con.ChargeSpeed
	connector.Type = con.Type
	connector.Active = con.Active

	return connector
}