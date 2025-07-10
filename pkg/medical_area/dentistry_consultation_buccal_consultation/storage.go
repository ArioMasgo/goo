package dentistry_consultation_buccal_consultation

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesBuccalConsultationRepository interface {
	create(m *BuccalConsultation) error
	update(m *BuccalConsultation) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*BuccalConsultation, error)
	getByIDConsultation(id string) (*BuccalConsultation, error)
	getAll() ([]*BuccalConsultation, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesBuccalConsultationRepository {
	return newBuccalConsultationSqlServerRepository(db, txID)
}
