package encuestas

type Encuesta struct {
	ID          int     `json:"id_encuesta" db:"id_encuesta"`
	Nombre      string  `json:"nombre_encuesta" db:"nombre_encuesta"`
	Descripcion *string `json:"descripcion,omitempty" db:"descripcion"`
	Estado      string  `json:"estado" db:"estado"`
	FechaInicio *string `json:"fecha_inicio,omitempty" db:"fecha_inicio"`
	FechaFin    *string `json:"fecha_fin,omitempty" db:"fecha_fin"`
}
