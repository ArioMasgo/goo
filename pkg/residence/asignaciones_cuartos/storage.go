package asignaciones_cuartos

import (
	"github.com/jmoiron/sqlx"

	"dbu-api/internal/models"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type RoomAssignmentRepository interface {
	create(m *RoomAssignment) error
	update(m *RoomAssignment) error
	delete(m *RoomAssignment) error
	getByID(id string) (*RoomAssignment, error)
	getAll() ([]*RoomAssignment, error)
	getRoomAssignmentByRoomIDSubmissionID(roomID string, submissionID int64) ([]*RoomAssignment, error)
	multiAssign(assignments []string) error
	getAllRoomAssignmentsByStudentIDANDSubmissionID(studentID, callID int64) ([]*RoomAssignment, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) RoomAssignmentRepository {
	return newAsignacionesCuartosSqlServerRepository(db, user, txID)
}
