package psicopedagogia

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type ResponseData struct {
	Error   bool   `json:"error"`
	Msg     string `json:"msg"`
	Detalle any    `json:"detalle"`
	Code    int    `json:"code"`
	Type    string `json:"type"`
}

var RequestRespuesta struct {
	IDParticipante int `json:"idParticipante"`
	NumeroAtencion int `json:"numeroAtencion"`
	Page           int `json:"page"`
	PageSize       int `json:"pageSize"`
}

type RequestCita struct {
	DNI         string    `json:"dni"`
	Nombre      string    `json:"nombre"`
	Apellido    string    `json:"apellido"`
	Facultad    string    `json:"facultad"`
	FechaInicio time.Time `json:"hora_inicio"`
	FechaFin    time.Time `json:"hora_fin"`
	Estado      string    `json:"estado"`
}

type RequestParticipantes struct {
	Tipo                 string  `json:"tipo_participante" valid:"required"`
	Nombre               string  `json:"nombre" valid:"required"`
	Apellido             string  `json:"apellido" valid:"required"`
	DNI                  string  `json:"dni"`
	NumTelefono          string  `json:"num_telefono"`
	Estado               string  `json:"estado"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
	ColegioProcedencia   string  `json:"colegio_procedencia"`
	AnioIngreso          int     `json:"anio_ingreso"`
	Escuela              string  `json:"escuela"`
	CodigoEstudiante     string  `json:"codigo_estudiante"`
	FechaNacimiento      string  `json:"fecha_nacimiento"`
	Edad                 int     `json:"edad"`
	LugarNacimiento      string  `json:"lugar_nacimiento"`
	ModalidadIngreso     string  `json:"modalidad_ingreso"`
	NumeroAtencion       int     `json:"numero_atencion"`
	Sexo                 string  `json:"sexo"`
	DiagnosticoID        int     `json:"diagnostico_id"`
	ConQuienesVive       string  `json:"con_quienes_vive_actualmente"`
	SemestreCursa        string  `json:"semestre_cursa"`
	Direccion            string  `json:"direccion"`
	QuienFinanciaCarrera string  `json:"quien_financia_carrera"`
	MotivoConsulta       string  `json:"motivo_consulta"`
	SituacionActual      string  `json:"situacion_actual"`
	OtrosProcedimientos  string  `json:"otros_procedimientos"`
	EsSRQ                bool    `json:"es_srq"`
	Notes                string  `json:"notes"`
	EstadoEvaluacion     string  `json:"estado_evaluacion"`
	Profesion            string  `json:"profesion"`
	EstadoCivil          *string `json:"estado_civil"`
	LaboraEnUnas         bool    `json:"labora_en_unas"`
	GradoInstruccion     *string `json:"grado_instruccion"`
	KeyUrl               string  `json:"key_url"`
}

type RequestPreguntas struct {
	TextoPregunta string `json:"texto_pregunta" valid:"required"`
	IsMandatory   bool   `json:"is_mandatory"`
	Order         int    `json:"order"`
	Type          string `json:"type" valid:"required"`
}

type RequestEncuestas struct {
	Nombre      string  `json:"nombre_encuesta" valid:"required"`
	Descripcion *string `json:"descripcion,omitempty"`
	Estado      string  `json:"estado" valid:"required"`
	FechaInicio *string `json:"fecha_inicio,omitempty"`
	FechaFin    *string `json:"fecha_fin,omitempty"`
}

type RequestRespuestas struct {
	IDParticipante int    `json:"id_participante" valid:"required"`
	IDEncuesta     int    `json:"id_encuesta" valid:"required"`
	IDPregunta     int    `json:"id_pregunta" valid:"required"`
	Respuesta      string `json:"respuesta" valid:"required"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type RespuestaData struct {
	IDPregunta int    `json:"id_pregunta"`
	Respuesta  string `json:"respuesta"`
}

type GuardarEncuestaRequest struct {
	Participante RequestParticipantes `json:"participante"`
	EncuestaID   int                  `json:"encuesta_id"`
	Respuestas   []RespuestaData      `json:"respuestas"`
}

type GuardarHistorialRequest struct {
	Participante RequestParticipantes `json:"participante"`
}

func (m *RequestRespuestas) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *RequestEncuestas) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *RequestPreguntas) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
