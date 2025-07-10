package rooms

import (
	"github.com/jmoiron/sqlx"

	"dbu-api/internal/models"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesRoomRepository interface {
	create(m *Room) error
	update(m *Room) error
	delete(id string) error
	getByID(id string) (*Room, error)
	getAll() ([]*Room, error)
	getAllRoomsByResidenceID(id string) ([]*Room, error)
	multiCreate(rooms []*Room) error
	updateOnlyCharacteristics(m *Room) error
	gtAllRoomsBySubmissionIDResidenceID(submissionID int64, residenceID string) ([]*Room, error)
	getRoomsByResidence(residenceID string, page, limit int) ([]*Room, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesRoomRepository {
	return newRoomSqlServerRepository(db, user, txID)
}
