package mocks

import (
	"github.com/sede-x/gopoc-connector/pkg/models"
)

var Connectors = []models.Connector{
	{
		Id:          1,
		StationId:   3,
		Type:        "L1, AC",
		ChargeSpeed: "2 kW",
		Active:      true,
	},
	{
		Id:          2,
		StationId:   3,
		Type:        "L2, AC",
		ChargeSpeed: "6 kW",
		Active:      true,
	},
	{
		Id:          3,
		StationId:   5,
		Type:        "DCFC", // Direct Current Fast Chargers (DCFC)
		ChargeSpeed: "50 KW",
		Active:      false,
	},
}
