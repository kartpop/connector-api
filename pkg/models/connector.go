package models

import "github.com/google/uuid"

type Connector struct {
	Id          uuid.UUID `json:"id"`
	StationId   uuid.UUID `json:"stationid"`
	Type        string    `json:"type"`
	ChargeSpeed string    `json:"chargespeed"`
	Active      bool      `json:"active"`
}
