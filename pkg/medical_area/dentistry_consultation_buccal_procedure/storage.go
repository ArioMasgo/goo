package dentistry_consultation_buccal_procedure

import (
	"dbu-api/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesBuccalProcedureRepository interface {
	create(m *BuccalProcedure) error
	update(m *BuccalProcedure) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*BuccalProcedure, error)
	getByIDConsultation(id string) (*BuccalProcedure, error)
	getAll() ([]*BuccalProcedure, error)
	existsByReceipt(recibo string) (bool, error)
	getBuccalProceduresExcel(fecha_inicio, fecha_fin string) ([]*models.PerformedProceduresExcel, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesBuccalProcedureRepository {
	return newBuccalProcedureSqlServerRepository(db, txID)
}
