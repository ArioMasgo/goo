package medical_general_medicine_consultation

import (
	"dbu-api/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesGeneralMedicineConsultationRepository interface {
	create(m *GeneralMedicineConsultation) error
	update(m *GeneralMedicineConsultation) error
	delete(id string) error
	getByID(id string) (*GeneralMedicineConsultation, error)
	getByIDPatient(id string) ([]*GeneralMedicineConsultation, error)
	getAll() ([]*GeneralMedicineConsultation, error)
	getAllByDateExcel(fecha_inicio, fecha_fin string) ([]*models.ConsultationIntegralAttentionExcel, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesGeneralMedicineConsultationRepository {
	return newGeneralMedicineConsultationSqlServerRepository(db, txID)
}
