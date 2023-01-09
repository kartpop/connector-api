package models

type ConnectorPagination struct {
	Limit      int           `json:"limit,omitempty"`
	Page       int           `json:"page,omitempty"`
	Sort       string        `json:"sort,omitempty"`
	TotalRows  int64         `json:"total_rows"`
	TotalPages int           `json:"total_pages"`
	Connectors *[]*Connector `json:"connectors"`
}

func (cp *ConnectorPagination) GetOffset() int {
	return (cp.GetPage() - 1) * cp.GetLimit()
}

func (cp *ConnectorPagination) GetLimit() int {
	if cp.Limit == 0 {
		cp.Limit = 30
	}
	return cp.Limit
}

func (cp *ConnectorPagination) GetPage() int {
	if cp.Page == 0 {
		cp.Page = 1
	}
	return cp.Page
}
