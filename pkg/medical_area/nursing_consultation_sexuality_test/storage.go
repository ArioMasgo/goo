package nursing_consultation_sexuality_test

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesSexualityTestRepository interface {
	create(m *SexualityTest) error
	update(m *SexualityTest) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*SexualityTest, error)
	getAll() ([]*SexualityTest, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesSexualityTestRepository {
	return newSexualityTestSqlServerRepository(db, txID)
}
