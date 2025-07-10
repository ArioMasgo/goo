package preguntas

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesPreguntaRepository interface {
	GetAll(page, pageSize int) ([]Pregunta, error)
	GetByID(id int) (*Pregunta, error)
	Create(p *Pregunta) error
	Update(id int, p *Pregunta) error
	Delete(id int) error
}

func FactoryStorage(db *sqlx.DB) ServicesPreguntaRepository {
	return NewPreguntaRepository(db)
}
