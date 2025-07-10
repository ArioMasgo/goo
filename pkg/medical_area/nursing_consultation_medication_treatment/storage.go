package nursing_consultation_medication_treatment

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesMedicationTreatmentRepository interface {
	create(m *MedicationTreatment) error
	update(m *MedicationTreatment) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) (*MedicationTreatment, error)
	getByIDConsultation(id string) ([]*MedicationTreatment, error)
	getAll() ([]*MedicationTreatment, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesMedicationTreatmentRepository {
	return newMedicationTreatmentSqlServerRepository(db, txID)
}
