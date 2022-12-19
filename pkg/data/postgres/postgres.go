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
	if result := pg.DB.Create(con); result.Error != nil {
		return result.Error
	}

	return nil
}

func (pg *PostgresDB) GetConnectorByID(id int) (*models.Connector, error) {
	var connector models.Connector
	if result := pg.DB.First(&connector, id); result.Error != nil {
		return nil, result.Error
	}

	return &connector, nil
}

func (pg *PostgresDB) UpdateConnector(id int, upcon models.Connector) (*models.Connector, error) {
	var con models.Connector
	if result := pg.DB.First(&con, id); result.Error != nil {
		return nil, result.Error
	}

	// TODO: check if updatedConnector has all fields set.
	// If it excludes some fields, default values would get set for the original connector in DB causing data loss.
	con.StationId = upcon.StationId
	con.Type = upcon.Type
	con.ChargeSpeed = upcon.ChargeSpeed
	con.Active = upcon.Active
	if result := pg.DB.Save(&con); result.Error != nil {
		return nil, result.Error
	}

	return &con, nil
}

func (pg *PostgresDB) DeleteConnector(id int) error {
	var connector models.Connector
	if result := pg.DB.First(&connector, id); result.Error != nil {
		return result.Error
	}
	if result := pg.DB.Delete(&connector); result.Error != nil {
		return result.Error
	}

	return nil
}
