package encuestas

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesEncuestaRepository interface {
	GetAll(page, pageSize int) ([]Encuesta, error)
	GetByID(id int) (*Encuesta, error)
	Create(e *Encuesta) (int64, error)
	Update(id int, e *Encuesta) error
	Delete(id int) error
	GetActiveSRQ() (int, string, bool, error)
}

func FactoryStorage(db *sqlx.DB) ServicesEncuestaRepository {
	return NewEncuestaRepository(db)
}
