package psicopedagogia

import (
	"dbu-api/internal/models"
	"dbu-api/pkg/psicopedagogia/alumnos"
	"dbu-api/pkg/psicopedagogia/citas"
	"dbu-api/pkg/psicopedagogia/diagnosticos"
	"dbu-api/pkg/psicopedagogia/encuestas"
	"dbu-api/pkg/psicopedagogia/gpdfs"
	"dbu-api/pkg/psicopedagogia/historialencuesta"
	"dbu-api/pkg/psicopedagogia/participantes"
	"dbu-api/pkg/psicopedagogia/preguntas"
	"dbu-api/pkg/psicopedagogia/respuestas"

	"github.com/jmoiron/sqlx"
)

type Serverpsicopedagogia struct {
	SrvEstudiantes       alumnos.PortsServerEstudiante
	SrvParticipantes     participantes.PortsServerParticipantes
	SrvPreguntas         preguntas.PortsServerPreguntas
	SrvEncuestas         encuestas.PortsServerEncuestas
	SrvRespuestas        respuestas.PortsServerRespuestas
	SrvHistorialEncuesta historialencuesta.PortsHistorialService
	SrvPDFs              gpdfs.PortsServerPDFs
	SrvDiagnostico       diagnosticos.PortsServerDiagnostico
	SrvCitas             citas.PortsServerCita
}

func NewServerpsicopedagogia(db *sqlx.DB, usr *models.User, txID string) *Serverpsicopedagogia {
	historialRepo := historialencuesta.FactoryStorage(db).(*historialencuesta.HistorialRepository)
	respuestaRepo := respuestas.FactoryStorage(db).(*respuestas.RespuestaRepository)
	participanteRepo := participantes.FactoryStorage(db).(*participantes.ParticipanteRepository)

	return &Serverpsicopedagogia{
		SrvEstudiantes:       alumnos.NewEstudianteService(alumnos.FactoryStorage(db, txID), txID),
		SrvParticipantes:     participantes.NewParticipanteService(participanteRepo, txID),
		SrvPreguntas:         preguntas.NewPreguntaService(preguntas.FactoryStorage(db), txID),
		SrvEncuestas:         encuestas.NewEncuestaService(encuestas.FactoryStorage(db), txID),
		SrvRespuestas:        respuestas.NewRespuestaService(respuestaRepo, txID),
		SrvHistorialEncuesta: historialencuesta.NewHistorialService(historialRepo, historialRepo, respuestaRepo, participanteRepo, txID),
		SrvPDFs:              gpdfs.NewPDFService(gpdfs.FactoryStorage(db), txID),
		SrvDiagnostico:       diagnosticos.NewDiagnosticoService(diagnosticos.FactoryStorage(db), txID),
		SrvCitas:             citas.NewCitaService(citas.FactoryStorage(db), txID),
	}
}
