package general_visit

import (
	"dbu-api/internal/models"
	"dbu-api/pkg/general_visit/visita_general"
	"dbu-api/pkg/general_visit/visita_residente"

	"github.com/jmoiron/sqlx"
)

type FactoryVisitaGeneral interface {
	VisitaGeneral() visita_general.PortsServerVisitaGeneral
	VisitaResidente() visita_residente.PortsServerVisitaResidente
}

type general struct {
	db   *sqlx.DB
	user *models.User
	txID string
}

func NewGeneralVisitService(db *sqlx.DB, user *models.User, txID string) FactoryVisitaGeneral {
	return &general{
		db:   db,
		user: user,
		txID: txID,
	}
}

func (s *general) VisitaGeneral() visita_general.PortsServerVisitaGeneral {
	repository := visita_general.FactoryStorage(s.db, s.txID)
	return visita_general.NewVisitaGeneralService(repository, s.user, s.txID)
}

func (s *general) VisitaResidente() visita_residente.PortsServerVisitaResidente {
	repository := visita_residente.FactoryStorage(s.db, s.txID)
	return visita_residente.NewVisitaResidenteService(repository, s.user, s.txID)
}