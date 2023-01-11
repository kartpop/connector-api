package models

import (
	"fmt"
	"time"

	"github.com/sede-x/gopoc-connector/pkg/helper"
)

type Connector struct {
	Id           string    `json:"id,omitempty" gorm:"primaryKey"`
	Name         string    `json:"name,omitempty" gorm:"unique"`
	CustomerId   string    `json:"customer_id,omitempty"`
	CustomerName string    `json:"customer_name,omitempty" gorm:"index:idx_custloctype,priority:1"`
	LocationId   string    `json:"location_id,omitempty"`
	LocationName string    `json:"location_name,omitempty" gorm:"index:idx_custloctype,priority:2"`
	Type         string    `json:"type,omitempty" gorm:"index:idx_custloctype,priority:3"`
	ChargeSpeed  string    `json:"charge_speed,omitempty"`
	Active       bool      `json:"active,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"autoUpdateTime"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (c Connector) GenerateId() string {
	keyString := fmt.Sprintf("%s-%s-%s", c.LocationId, c.Type, c.ChargeSpeed)
	return helper.GetMD5Hash(keyString)
}
