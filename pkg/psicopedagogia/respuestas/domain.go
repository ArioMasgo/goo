package respuestas

type Respuesta struct {
	ID             int    `json:"id_respuesta" db:"id_respuesta"`
	IDParticipante int    `json:"id_participante" db:"id_participante"`
	IDEncuesta     int    `json:"id_encuesta" db:"id_encuesta"`
	IDPregunta     int    `json:"id_pregunta" db:"id_pregunta"`
	IdHistorial    int    `json:"id_historial" db:"id_historial"`
	Respuesta      string `json:"respuesta" db:"respuesta"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	UpdatedAt      string `json:"updated_at" db:"updated_at"`
	NumeroAtencion int    `json:"numero_atencion" db:"numero_atencion"`
}

type RespuestaDetalle struct {
	IDRespuesta    int    `json:"id_respuesta" db:"id_respuesta"`
	IDParticipante int    `json:"id_participante" db:"id_participante"`
	IDEncuesta     int    `json:"id_encuesta" db:"id_encuesta"`
	IDPregunta     int    `json:"id_pregunta" db:"id_pregunta"`
	TextoPregunta  string `json:"texto_pregunta" db:"texto_pregunta"`
	NumeroAtencion int    `json:"numero_atencion" db:"numero_atencion"`
	Respuesta      string `json:"respuesta" db:"respuesta"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	UpdatedAt      string `json:"updated_at" db:"updated_at"`
}
