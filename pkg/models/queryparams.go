package models

type ConnectorQueryParams struct {
	LocationIds []string
	Types       []string
	Sort        bool // if true { sort by 1.location_id, 2.type } else { don't sort }
	Limit       int
	Page        int
}
