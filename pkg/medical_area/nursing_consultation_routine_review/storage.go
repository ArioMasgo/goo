package nursing_consultation_routine_review

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesRoutineReviewRepository interface {
	create(m *RoutineReview) error
	update(m *RoutineReview) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*RoutineReview, error)
	getAll() ([]*RoutineReview, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesRoutineReviewRepository {
	return newRoutineReviewSqlServerRepository(db, txID)
}
