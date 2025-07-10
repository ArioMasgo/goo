package dentistry_consultation_odontogram_review

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesOdontogramReviewRepository interface {
	create(m *OdontogramReview) error
	update(m *OdontogramReview) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*OdontogramReview, error)
	getByIDConsultation(id string) (*OdontogramReview, error)
	getAll() ([]*OdontogramReview, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesOdontogramReviewRepository {
	return newOdontogramReviewSqlServerRepository(db, txID)
}
