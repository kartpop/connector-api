package models

import (
	"fmt"
	"time"

	"github.com/sede-x/gopoc-connector/pkg/helper"
)

type Connector struct {
	Id          string    `json:"id,omitempty" gorm:"primaryKey"`
	LocationId  string    `json:"location_id,omitempty" gorm:"index:idx_locidtype,priority:1"`
	Type        string    `json:"type,omitempty" gorm:"index:idx_locidtype,priority:2"`
	ChargeSpeed string    `json:"charge_speed,omitempty"`
	Active      bool      `json:"active,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"autoUpdateTime"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (c Connector) GenerateId() string {
	keyString := fmt.Sprintf("%s-%s-%s", c.LocationId, c.Type, c.ChargeSpeed)
	return helper.GetMD5Hash(keyString)
}
