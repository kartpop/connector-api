package mocks

import (
	"github.com/kartpop/connector-api/pkg/helper"
	"github.com/kartpop/connector-api/pkg/models"
)

var Connectors = []*models.Connector{
	{
		Id:          helper.GetMD5Hash("lsdfy232"),
		LocationId:  helper.GetMD5Hash("lsfddffy23432"),
		Type:        "L1, AC",
		ChargeSpeed: "2 kW",
		Active:      true,
	},
	{
		Id:          helper.GetMD5Hash("sdf544"),
		LocationId:  helper.GetMD5Hash("sdfd787"),
		Type:        "L2, AC",
		ChargeSpeed: "6 kW",
		Active:      true,
	},
	{
		Id:          helper.GetMD5Hash("2343sdfd"),
		LocationId:  helper.GetMD5Hash("656sdfd"),
		Type:        "L3",
		ChargeSpeed: "50 KW",
		Active:      false,
	},
}
