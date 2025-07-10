package consultation_medical_area

import (
	"dbu-api/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesConsultationMedicalAreaRepository interface {
	create(m *ConsultationMedicalArea) error
	update(m *ConsultationMedicalArea) error
	delete(id string) error
	getByID(id string) (*ConsultationMedicalArea, error)
	getAll() ([]*ConsultationMedicalArea, error)
	GetAllByPatientID(id string) ([]*ConsultationMedicalArea, error)
	getByPatientDNI(dni string) ([]*ConsultationMedicalArea, error)
	getAllByNursingDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error)
	getAllByDentistryDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error)
	getAllByMedicalDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesConsultationMedicalAreaRepository {
	return newConsultationMedicalAreaSqlServerRepository(db, txID)
}
