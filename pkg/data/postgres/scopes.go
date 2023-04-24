package postgres

import (
	"math"

	"github.com/kartpop/connector-api/pkg/models"
	"gorm.io/gorm"
)

func Paginate(connectors *[]*models.Connector, cp *models.ConnectorPagination, qp models.ConnectorQueryParams, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	cp.Limit = qp.Limit
	cp.Page = qp.Page

	var totalRows int64
	db.Model(connectors).Count(&totalRows)
	cp.TotalRows = totalRows

	totalPages := int(math.Ceil(float64(totalRows) / float64(cp.GetLimit())))
	cp.TotalPages = totalPages

	if qp.Sort {
		cp.Sort = "customer_name, location_name, type"
		return func(db *gorm.DB) *gorm.DB {
			return db.Offset(cp.GetOffset()).Limit(cp.GetLimit()).Order("customer_name").Order("location_name").Order("type")
		}
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(cp.GetOffset()).Limit(cp.GetLimit())
	}
}
