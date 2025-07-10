package patients

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesPatientsRepository interface {
	create(m *Patients) error
	update(m *Patients) error
	delete(id string) error
	getByID(id string) (*Patients, error)
	getByDNI(dni string) (*Patients, error)
	getAll() ([]*Patients, error)
	countPaginationPatients(dni, names, surnames string) (int64, error)
	searchPaginationPatients(dni, names, surnames string, limit, offset int64) ([]*Patients, error)
	existsByDNI(dni string) (bool, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesPatientsRepository {
	return newPatientsSqlServerRepository(db, txID)
}
