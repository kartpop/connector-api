package postgres

import (
	"github.com/sede-x/gopoc-connector/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func New(dbURL string) (*PostgresDB, error) {
	gormdb, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = gormdb.AutoMigrate(&models.Connector{})
	if err != nil {
		return nil, err
	}

	return &PostgresDB{gormdb}, nil
}

func (pg *PostgresDB) GetAllConnectors() (*[]models.Connector, error) {
	var connectors []models.Connector
	if result := pg.DB.Find(&connectors); result.Error != nil {
		return nil, result.Error
	}

	return &connectors, nil
}

func (pg *PostgresDB) AddConnector(con *models.Connector) error {
	con.Id = con.GenerateId()
	if result := pg.DB.Create(con); result.Error != nil {
		return result.Error
	}

	return nil
}

func (pg *PostgresDB) GetConnectorByID(id string) (*models.Connector, error) {
	var con models.Connector
	if result := pg.DB.First(&con, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}

	return &con, nil
}

func (pg *PostgresDB) UpdateConnector(id string, upcon models.Connector) (*models.Connector, error) {
	var con models.Connector
	if result := pg.DB.First(&con, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}

	// TODO: check if updatedConnector has all fields set.
	// If it excludes some fields, default values would get set for the original connector in DB causing data loss.
	con.LocationId = upcon.LocationId
	con.Type = upcon.Type
	con.ChargeSpeed = upcon.ChargeSpeed
	con.Active = upcon.Active
	if result := pg.DB.Save(&con); result.Error != nil {
		return nil, result.Error
	}

	return &con, nil
}

func (pg *PostgresDB) DeleteConnector(id string) error {
	var con models.Connector
	if result := pg.DB.First(&con, "id = ?", id); result.Error != nil {
		return result.Error
	}
	if result := pg.DB.Delete(&con); result.Error != nil {
		return result.Error
	}

	return nil
}

func (pg *PostgresDB) GetConnectors(qp models.ConnectorQueryParams) ([]*models.Connector, error) {
	var connectors []*models.Connector
	result := pg.DB

	// CAUTION: Take care while chaining methods, refer - https://gorm.io/docs/method_chaining.html
	// use sorting if required
	if qp.Sort {
		result = result.Order("location_id").Order("type")
	}

	// fetch data 
	if len(qp.LocationIds) > 0 && len(qp.Types) > 0 {
		result = result.Where("location_id IN ? AND type IN ?", qp.LocationIds, qp.Types).Find(&connectors)
	} else if len(qp.LocationIds) > 0 {
		result = result.Where("location_id IN ?", qp.LocationIds).Find(&connectors)
	} else if len(qp.Types) > 0 {
		result = result.Where("type IN ?", qp.Types).Find(&connectors)
	} else { // get all connectors
		result = result.Find(&connectors)
	}
	
	if result.Error != nil {
		return nil, result.Error
	}

	return connectors, nil
}

