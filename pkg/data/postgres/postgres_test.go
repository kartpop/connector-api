package postgres

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sede-x/gopoc-connector/pkg/helper"
	"github.com/sede-x/gopoc-connector/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupPostgresDB(t *testing.T) (*PostgresDB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error: '%s' was not expected while creating sqlmock", err.Error())
	}

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Errorf("error: '%s' was not expected while opening gorm DB", err.Error())
	}

	return &PostgresDB{DB: gdb}, mock, func() {
		db.Close()
	}
}

func TestGetConnectorByIDValidRequest(t *testing.T) {
	// setup
	pg, mock, close := SetupPostgresDB(t)
	defer close()
	var (
		id   = helper.GetMD5Hash("test-id")
		name = "test-con-1"
	)
	query := "SELECT * FROM \"connectors\" WHERE id = $1 ORDER BY \"connectors\".\"id\" LIMIT 1"
	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(id).WillReturnRows(rows)

	// invoke
	con, err := pg.GetConnectorByID(id)

	// test
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(&models.Connector{Id: id, Name: name}, con) {
		t.Errorf("Incorrect connector retrieved from DB")
	}
}

func TestGetConnectorByIDErrNotFound(t *testing.T) {
	// setup
	pg, mock, close := SetupPostgresDB(t)
	defer close()
	query := "SELECT * FROM \"connectors\" WHERE id = $1 ORDER BY \"connectors\".\"id\" LIMIT 1"
	rows := sqlmock.NewRows([]string{"id", "name"})
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("any-id").WillReturnRows(rows)

	// invoke
	_, err := pg.GetConnectorByID("any-id")

	// test
	if err != nil {
		switch e := err.(type) {
		case *helper.ErrRecordNotFound:
		default:
			t.Errorf("Expected error 'record not found', got: '%s'", e.Error())
		}
	} else {
		t.Errorf("Expected error 'record not found', got no error.")
	}
}
