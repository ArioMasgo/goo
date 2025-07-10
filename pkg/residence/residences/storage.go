package residences

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServiceResidenceRepository interface {
	create(m *Residence) error
	update(m *Residence) error
	delete(id string) error
	getByID(id string) (*Residence, error)
	getAll() ([]*Residence, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServiceResidenceRepository {
	return newResidenceSqlServerRepository(db, txID)
}
