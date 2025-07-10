package psicopedagogia

import (
	"dbu-api/pkg/psicopedagogia/historialencuesta"
	"dbu-api/pkg/psicopedagogia/participantes"
	"dbu-api/pkg/psicopedagogia/respuestas"
	"time"
)

func ConvertRequestToParticipante(r *RequestParticipantes) *participantes.Participante {
	return &participantes.Participante{
		Tipo:               r.Tipo,
		Nombre:             r.Nombre,
		Apellido:           r.Apellido,
		DNI:                r.DNI,
		Estado:             r.Estado,
		CreatedAt:          time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:          time.Now().Format("2006-01-02 15:04:05"),
		ColegioProcedencia: r.ColegioProcedencia,
		AnioIngreso:        r.AnioIngreso,
		Escuela:            r.Escuela,
		CodigoEstudiante:   r.CodigoEstudiante,
		FechaNacimiento:    r.FechaNacimiento,
		Edad:               r.Edad,
		LugarNacimiento:    r.LugarNacimiento,
		ModalidadIngreso:   r.ModalidadIngreso,
		NumeroAtencion:     r.NumeroAtencion,
		Sexo:               r.Sexo,
		NumTelefono:        r.NumTelefono,
		ConQuienesVive:     r.ConQuienesVive,
		SemestreCursa:      r.SemestreCursa,
		Direccion:          r.Direccion,
		Profesion:          r.Profesion,
		EstadoCivil:        r.EstadoCivil,
		LaboraEnUnas:       r.LaboraEnUnas,
		GradoInstruccion:   r.GradoInstruccion,
	}
}

func ConvertRequestToRespuestas(respuestasData []RespuestaData, encuestaID int) []respuestas.Respuesta {
	var respuestasList []respuestas.Respuesta

	for _, r := range respuestasData {
		respuestasList = append(respuestasList, respuestas.Respuesta{
			IDParticipante: 0,
			IDEncuesta:     encuestaID,
			IDPregunta:     r.IDPregunta,
			Respuesta:      r.Respuesta,
			CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:      time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return respuestasList
}

func ConvertRequestToHistorialEncuesta(r *RequestParticipantes, idEncuesta int) *historialencuesta.HistorialEncuesta {
	return &historialencuesta.HistorialEncuesta{
		IDEncuesta:           idEncuesta,
		FechaRespuesta:       time.Now(),
		NumTelefono:          r.NumTelefono,
		ConQuienesVive:       r.ConQuienesVive,
		EstadoEvaluacion:     r.Estado,
		SemestreCursa:        r.SemestreCursa,
		Direccion:            r.Direccion,
		QuienFinanciaCarrera: r.QuienFinanciaCarrera,
		MotivoConsulta:       r.MotivoConsulta,
		SituacionActual:      r.SituacionActual,
		OtrosProcedimientos:  r.OtrosProcedimientos,
		DiagnosticoID:        r.DiagnosticoID,
		CreatedDate:          r.CreatedAt,
		EsSRQ:                r.EsSRQ,
		KeyUrl:               r.KeyUrl,
	}
}
func ConvertRequestToHistorial(r *RequestParticipantes) *historialencuesta.Historial {
	return &historialencuesta.Historial{
		DNI:                  r.DNI,
		FechaRespuesta:       time.Now(),
		NumTelefono:          r.NumTelefono,
		ConQuienesVive:       r.ConQuienesVive,
		EstadoEvaluacion:     r.Estado,
		SemestreCursa:        r.SemestreCursa,
		Direccion:            r.Direccion,
		QuienFinanciaCarrera: r.QuienFinanciaCarrera,
		MotivoConsulta:       r.MotivoConsulta,
		SituacionActual:      r.SituacionActual,
		OtrosProcedimientos:  r.OtrosProcedimientos,
		DiagnosticoID:        r.DiagnosticoID,
		EsSRQ:                r.EsSRQ,
	}
}
