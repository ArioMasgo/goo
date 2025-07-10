package participantes

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesParticipanteRepository interface {
	GetAll(page, pageSize int) ([]Participante, int, error)
	GetByID(id int) (*Participante, error)
	Create(p *Participante) (int, error)
	Update(id int, p *Participante) error
	Delete(id int) error
	SearchParticipants(criteria map[string]interface{}) ([]Participante, error)
}

func FactoryStorage(db *sqlx.DB) ServicesParticipanteRepository {
	return NewParticipanteRepository(db)
}
