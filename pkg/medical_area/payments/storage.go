package payments_concept

import (
	"github.com/jmoiron/sqlx"
)

type ServicesPaymentConceptRepository interface {
	search(area, tipoServicio, nombreServicio string) (*ServicioMedicoConfig, error)
	searchPaymentProcedureOdontologia(recibo, servicio string) ([]*PagosServicios, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesPaymentConceptRepository {
	return newPaymentConceptSqlServerRepository(db, txID)
}
