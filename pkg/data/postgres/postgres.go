package postgres

import (
	"errors"

	"github.com/kartpop/connector-api/pkg/helper"
	"github.com/kartpop/connector-api/pkg/models"
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

func (pg *PostgresDB) GetConnectors(qp models.ConnectorQueryParams) (*models.ConnectorPagination, error) {
	var pagedConnectors models.ConnectorPagination
	result := FilterConnectors(pg.DB, qp, &pagedConnectors)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pagedConnectors, nil
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
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &helper.ErrRecordNotFound{}
		}
		return nil, result.Error
	}

	return &con, nil
}

func (pg *PostgresDB) UpdateConnector(id string, upcon models.Connector) (*models.Connector, error) {
	var con models.Connector
	if result := pg.DB.First(&con, "id = ?", id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &helper.ErrRecordNotFound{}
		}
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
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &helper.ErrRecordNotFound{}
		}
		return result.Error
	}
	if result := pg.DB.Delete(&con); result.Error != nil {
		return result.Error
	}

	return nil
}
