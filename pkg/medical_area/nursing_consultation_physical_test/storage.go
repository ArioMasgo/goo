package nursing_consultation_physical_test

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesPhysicalTestRepository interface {
	create(m *PhysicalTest) error
	update(m *PhysicalTest) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*PhysicalTest, error)
	getAll() ([]*PhysicalTest, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesPhysicalTestRepository {
	return newPhysicalTestSqlServerRepository(db, txID)
}
