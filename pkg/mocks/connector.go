package mocks

import (
	"github.com/google/uuid"
	"github.com/sede-x/gopoc-connector/pkg/models"
)

var Connectors = []models.Connector{
	{
		Id:          uuid.New(),
		StationId:   uuid.New(),
		Type:        "L1, AC",
		ChargeSpeed: "2 kW",
		Active:      true,
	},
	{
		Id:          uuid.New(),
		StationId:   uuid.New(),
		Type:        "L2, AC",
		ChargeSpeed: "6 kW",
		Active:      true,
	},
	{
		Id:          uuid.New(),
		StationId:   uuid.New(),
		Type:        "DCFC", // Direct Current Fast Chargers (DCFC)
		ChargeSpeed: "50 KW",
		Active:      false,
	},
}
