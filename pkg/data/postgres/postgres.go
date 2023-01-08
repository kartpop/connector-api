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

func (pg *PostgresDB) GetConnectors(locationIds []string, types []string) (*[]models.Connector, error) {
	var connectors []models.Connector
	var result *gorm.DB
	if len(locationIds) > 0 && len(types) > 0 {
		result = pg.DB.Where("location_id IN ? AND type IN ?", locationIds, types).Find(&connectors)
	} else if len(locationIds) > 0 {
		result = pg.DB.Where("location_id IN ?", locationIds).Find(&connectors)
	} else if len(types) > 0 {
		result = pg.DB.Where("type IN ?", types).Find(&connectors)
	} else {
		result = pg.DB.Find(&connectors)
	}
	
	if result.Error != nil {
		return nil, result.Error
	}

	return &connectors, nil
}
