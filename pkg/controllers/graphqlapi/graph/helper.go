package graph

import (
	"github.com/sede-x/gopoc-connector/pkg/controllers/graphqlapi/graph/model"
	"github.com/sede-x/gopoc-connector/pkg/models"
)

func convertToGraphConnector(con models.Connector) model.Connector {
	var connector model.Connector
	connector.ID = con.Id
	connector.LocationID = con.LocationId
	connector.ChargeSpeed = con.ChargeSpeed
	connector.Type = con.Type
	connector.Active = con.Active

	return connector
}

func convertToDBConnector(con model.Connector) models.Connector {
	var connector models.Connector
	if con.ID != "" {
		connector.Id = con.ID
	}
	connector.LocationId = con.LocationID
	connector.ChargeSpeed = con.ChargeSpeed
	connector.Type = con.Type
	connector.Active = con.Active

	return connector
}
