package dentistry_consultation_buccal_test

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesBuccalTestRepository interface {
	create(m *BuccalTest) error
	update(m *BuccalTest) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*BuccalTest, error)
	getByIDConsultation(id string) (*BuccalTest, error)
	getAll() ([]*BuccalTest, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesBuccalTestRepository {
	return newBuccalTestSqlServerRepository(db, txID)
}
