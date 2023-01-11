package models

type ConnectorQueryParams struct {
	CustomerNames []string `json:"customer_names,omitempty"`
	LocationNames []string `json:"location_names,omitempty"`
	Types         []string `json:"types,omitempty"`
	Sort          bool `json:"sort,omitempty"` // if true { sort by 1.customer_name, 2.location_name, 3.type } else { don't sort }
	Limit         int `json:"limit,omitempty"`
	Page          int `json:"page,omitempty"`
}
