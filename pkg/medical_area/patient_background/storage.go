package patient_background

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesPatientBackgroundRepository interface {
	create(m *PatientBackground) error
	update(m *PatientBackground) error
	delete(id string) error
	deleteByIDPatient(id string) error
	getByID(id string) (*PatientBackground, error)
	getByIDPatient(id string) ([]*PatientBackground, error)
	getAll() ([]*PatientBackground, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesPatientBackgroundRepository {
	return newPatientBackgroundSqlServerRepository(db, txID)
}
