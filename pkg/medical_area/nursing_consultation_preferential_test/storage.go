package nursing_consultation_preferential_test

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesPreferentialTestRepository interface {
	create(m *PreferentialTest) error
	update(m *PreferentialTest) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*PreferentialTest, error)
	getAll() ([]*PreferentialTest, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesPreferentialTestRepository {
	return newPreferentialTestSqlServerRepository(db, txID)
}
