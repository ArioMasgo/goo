package residence_configuration

import (
	"github.com/jmoiron/sqlx"
)

type ServiceResidenceConfigurationRepository interface {
	create(m *ResidenceConfiguration) error
	update(m *ResidenceConfiguration) error
	delete(id string) error
	getByID(id string) (*ResidenceConfiguration, error)
	getAll() ([]*ResidenceConfiguration, error)
	getByResidenceID(residenciaID string) (*ResidenceConfiguration, error)
	updateByResidenceID(m *ResidenceConfiguration) error
}

func FactoryStorage(db *sqlx.DB, txID string) ServiceResidenceConfigurationRepository {
	return newResidenceConfigurationSqlServerRepository(db, txID)
}
