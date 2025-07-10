package preguntas

import (
	"github.com/jmoiron/sqlx"
)

type PreguntaRepository struct {
	DB *sqlx.DB
}

func NewPreguntaRepository(db *sqlx.DB) *PreguntaRepository {
	return &PreguntaRepository{DB: db}
}

func (r *PreguntaRepository) GetAll(page, pageSize int) ([]Pregunta, error) {
	var preguntas []Pregunta
	offset := (page - 1) * pageSize
	query := "SELECT * FROM questions LIMIT ? OFFSET ?"
	err := r.DB.Select(&preguntas, query, pageSize, offset)
	return preguntas, err
}

func (r *PreguntaRepository) GetByID(id int) (*Pregunta, error) {
	var pregunta Pregunta
	err := r.DB.Get(&pregunta, "SELECT * FROM questions WHERE id = ?", id)
	return &pregunta, err
}

func (r *PreguntaRepository) Create(p *Pregunta) error {
	_, err := r.DB.NamedExec(`
		INSERT INTO questions (texto_pregunta, is_mandatory, `+"`order`"+`, type)
		VALUES (:texto_pregunta, :is_mandatory, :order, :type)`, p)
	return err
}

func (r *PreguntaRepository) Update(id int, p *Pregunta) error {
	query := "UPDATE questions SET "
	params := make(map[string]interface{})

	if p.TextoPregunta != "" {
		query += "texto_pregunta=:texto_pregunta, "
		params["texto_pregunta"] = p.TextoPregunta
	}
	if p.IsMandatory {
		query += "is_mandatory=:is_mandatory, "
		params["is_mandatory"] = p.IsMandatory
	}
	if p.Order != 0 {
		query += "`order`=:order, "
		params["order"] = p.Order
	}
	if p.Type != "" {
		query += "type=:type, "
		params["type"] = p.Type
	}

	query = query[:len(query)-2] + " WHERE id=:id"
	params["id"] = id

	_, err := r.DB.NamedExec(query, params)
	return err
}

func (r *PreguntaRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM questions WHERE id = ?", id)
	return err
}
