package models

import (
	"fmt"
	"time"

	"github.com/sede-x/gopoc-connector/pkg/helper"
)

type Connector struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	LocationId  string    `json:"locationid" gorm:"index:idx_locidtype,priority:1"`
	Type        string    `json:"type" gorm:"index:idx_locidtype,priority:2"`
	ChargeSpeed string    `json:"chargespeed"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"createdat" gorm:"autoUpdateTime"`
	UpdatedAt   time.Time `json:"updatedat" gorm:"autoUpdateTime"`
}

func (c Connector) GenerateId() string {
	keyString := fmt.Sprintf("%s-%s-%s", c.LocationId, c.Type, c.ChargeSpeed)
	return helper.GetMD5Hash(keyString)
}

