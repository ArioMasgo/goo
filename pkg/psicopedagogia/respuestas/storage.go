package respuestas

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesRespuestaRepository interface {
	GetAll(page, pageSize int) ([]Respuesta, error)
	GetByID(id int) (*Respuesta, error)
	Create(res *Respuesta) error
	Update(id int, res *Respuesta) error
	Delete(id int) error
	GetResponsesPerParticipant(idParticipante int) ([]RespuestaDetalle, error)
	GetAllByParticipanteIdAndNumeroAtencion(participanteId, numeroAtencion, page, pageSize int) ([]RespuestaDetalle, error)
}

func FactoryStorage(db *sqlx.DB) ServicesRespuestaRepository {
	return NewRespuestaRepository(db)
}
