package models

type ConnectorQueryParams struct {
	LocationNames []string
	CustomerNames []string
	Types         []string
	Sort          bool // if true { sort by 1.customer_name, 2.location_name, 3.type } else { don't sort }
	Limit         int
	Page          int
}
