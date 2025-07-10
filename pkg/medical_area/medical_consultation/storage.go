package medical_consultation

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesMedicalConsultationRepository interface {
	create(m *MedicalConsultation) error
	update(m *MedicalConsultation) error
	delete(id string) error
	getByID(id string) (*MedicalConsultation, error)
	getByIDPatient(id string) ([]*MedicalConsultation, error)
	getAll() ([]*MedicalConsultation, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesMedicalConsultationRepository {
	return newMedicalConsultationSqlServerRepository(db, txID)
}
