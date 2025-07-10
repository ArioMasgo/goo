package convocatorias

import (
	"github.com/jmoiron/sqlx"

	"dbu-api/internal/models"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesConvocatoriasRepository interface {
	create(m *Convocatorias) error
	update(m *Convocatorias) error
	delete(id int64) error
	getByID(id int64) (*Convocatorias, error)
	getAll() ([]*Convocatorias, error)
	getAllByService(id int64) ([]*Convocatorias, error)
	getActive() (*Convocatorias, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesConvocatoriasRepository {
	return newConvocatoriasSqlServerRepository(db, user, txID)
}
