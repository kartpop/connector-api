package models

type Connector struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	StationId   int    `json:"stationid"`
	Type        string `json:"type"`
	ChargeSpeed string `json:"chargespeed"`
	Active      bool   `json:"active"`
}
