package dentistry_consultation

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesDentistryConsultationRepository interface {
	create(m *DentistryConsultation) error
	update(m *DentistryConsultation) error
	delete(id string) error
	getByID(id string) (*DentistryConsultation, error)
	getAll() ([]*DentistryConsultation, error)
	getByIDPatient(id string) ([]*DentistryConsultation, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesDentistryConsultationRepository {
	return newDentistryConsultationSqlServerRepository(db, txID)
}
