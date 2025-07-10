package residence_robot

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServiceResidenceRobotRepository interface {
	create(m *ResidenceRobot) error
	update(m *ResidenceRobot) error
	getByID(id int64) (*ResidenceRobot, error)
	getAll() ([]*ResidenceRobot, error)
	getByResidenceID(residenciaID string) (*ResidenceRobot, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServiceResidenceRobotRepository {
	return newResidenceRobotSqlServerRepository(db, txID)
}
