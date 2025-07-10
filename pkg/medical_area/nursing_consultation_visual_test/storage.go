package nursing_consultation_visual_test

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesVisualTestRepository interface {
	create(m *VisualTest) error
	update(m *VisualTest) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*VisualTest, error)
	getAll() ([]*VisualTest, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesVisualTestRepository {
	return newVisualTestSqlServerRepository(db, txID)
}
