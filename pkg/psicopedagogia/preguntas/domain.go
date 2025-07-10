package preguntas

type Pregunta struct {
	ID            int    `json:"id" db:"id"`
	TextoPregunta string `json:"texto_pregunta" db:"texto_pregunta"`
	IsMandatory   bool   `json:"is_mandatory" db:"is_mandatory"`
	Order         int    `json:"order" db:"order"`
	Type          string `json:"type" db:"type"`
}
