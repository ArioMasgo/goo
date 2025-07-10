package nursing_consultation_accompanying_data

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesAccompanyingDataRepository interface {
	create(m *AccompanyingData) error
	update(m *AccompanyingData) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*AccompanyingData, error)
	getAll() ([]*AccompanyingData, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesAccompanyingDataRepository {
	return newAccompanyingDataSqlServerRepository(db, txID)
}
