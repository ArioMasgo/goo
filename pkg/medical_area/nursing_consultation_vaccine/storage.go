package nursing_consultation_vaccine

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesVaccineRepository interface {
	create(m *Vaccine) error
	update(m *Vaccine) error
	delete(id string) error
	deleteByIDConsultation(id string) error
	getByID(id string) ([]*Vaccine, error)
	getByIDPatient(id string) ([]*Vaccine, error)
	getAll() ([]*Vaccine, error)
	getAllTypesVaccines() ([]*TypesVaccines, error)
	getAllVaccinesByPatientDni(dni string) ([]*Vaccine, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesVaccineRepository {
	return newVaccineSqlServerRepository(db, txID)
}
