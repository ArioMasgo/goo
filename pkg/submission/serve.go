package submission

import (
	"dbu-api/internal/models"
	"dbu-api/pkg/submission/alumnos"
	"dbu-api/pkg/submission/convocatorias"
	"github.com/jmoiron/sqlx"
)

type ServerResidence struct {
	SrvConvocatorias convocatorias.PortsServerConvocatorias
	SrvAlumnos       alumnos.PortsServerAlumnos
}

func NewServerSubmission(db *sqlx.DB, usr *models.User, txID string) *ServerResidence {
	return &ServerResidence{
		SrvConvocatorias: convocatorias.NewConvocatoriasService(convocatorias.FactoryStorage(db, usr, txID), usr, txID),
		SrvAlumnos:       alumnos.NewAlumnosService(alumnos.FactoryStorage(db, usr, txID), usr, txID),
	}
}
