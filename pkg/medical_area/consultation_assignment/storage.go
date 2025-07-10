package consultation_assignment

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesConsultationAssignmentRepository interface {
	create(m *ConsultationAssignment) error
	update(m *ConsultationAssignment) error
	updateByIDConsultation(m *ConsultationAssignment) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*ConsultationAssignment, error)
	getByIDConsultation(id string) (*ConsultationAssignment, error)
	getAll() ([]*ConsultationAssignment, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesConsultationAssignmentRepository {
	return newConsultationAssignmentSqlServerRepository(db, txID)
}
