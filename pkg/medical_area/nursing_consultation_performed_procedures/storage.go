package nursing_consultation_performed_procedures

import (
	"dbu-api/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesPerformedProceduresRepository interface {
	create(m *PerformedProcedures) error
	update(m *PerformedProcedures) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*PerformedProcedures, error)
	getAll() ([]*PerformedProcedures, error)
	GetByIDConsultation(id string) ([]*PerformedProcedures, error)
	getAllByDateExcel(fecha_inicio, fecha_fin string) ([]*models.PerformedProceduresExcel, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesPerformedProceduresRepository {
	return newPerformedProceduresSqlServerRepository(db, txID)
}
