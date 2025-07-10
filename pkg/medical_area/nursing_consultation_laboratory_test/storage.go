package nursing_consultation_laboratory_test

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesLaboratoryTestRepository interface {
	create(m *LaboratoryTest) error
	update(m *LaboratoryTest) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*LaboratoryTest, error)
	getAll() ([]*LaboratoryTest, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesLaboratoryTestRepository {
	return newLaboratoryTestSqlServerRepository(db, txID)
}
