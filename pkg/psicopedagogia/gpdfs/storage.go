package gpdfs

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesGpdfsRepository interface {
	GetResponsesPerParticipant(idParticipante, numeroAtencion int) ([]DataPDFRespuesta, error)
	GetParticipantByID(idParticipante int) (*ParticipantePDF, error)
}

func FactoryStorage(db *sqlx.DB) ServicesGpdfsRepository {
	return NewGpdfsRepository(db)
}
