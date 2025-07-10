package consultation_integral_attention

import (
	"dbu-api/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesConsultationIntegralAttentionRepository interface {
	create(m *ConsultationIntegralAttention) error
	update(m *ConsultationIntegralAttention) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*ConsultationIntegralAttention, error)
	getByIDConsultation(id string) (*ConsultationIntegralAttention, error)
	getAll() ([]*ConsultationIntegralAttention, error)
	getAllByDateExcel(fecha_inicio, fecha_fin string) ([]*models.ConsultationIntegralAttentionExcel, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesConsultationIntegralAttentionRepository {
	return newConsultationIntegralAttentionSqlServerRepository(db, txID)
}
