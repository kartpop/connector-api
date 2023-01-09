package postgres

import (
	"github.com/sede-x/gopoc-connector/pkg/models"
	"gorm.io/gorm"
)

func FilterConnectors(db *gorm.DB, qp models.ConnectorQueryParams, pagedConnectors *models.ConnectorPagination) *gorm.DB {
	var connectors []*models.Connector

	// CAUTION: Take care while chaining methods, refer - https://gorm.io/docs/method_chaining.html
	result := db.Session(&gorm.Session{})

	// filter
	if len(qp.LocationIds) > 0 && len(qp.Types) > 0 {
		result = result.Where("location_id IN ? AND type IN ?", qp.LocationIds, qp.Types)
	} else if len(qp.LocationIds) > 0 {
		result = result.Where("location_id IN ?", qp.LocationIds)
	} else if len(qp.Types) > 0 {
		result = result.Where("type IN ?", qp.Types)
	}

	// paginate and sort
	result = result.Scopes(Paginate(&connectors, pagedConnectors, qp, result)).Find(&connectors)
	pagedConnectors.Connectors = &connectors

	return result
}
