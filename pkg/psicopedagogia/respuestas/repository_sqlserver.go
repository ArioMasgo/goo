package respuestas

import (
	"github.com/jmoiron/sqlx"
)

type RespuestaRepository struct {
	DB *sqlx.DB
}

func NewRespuestaRepository(db *sqlx.DB) *RespuestaRepository {
	return &RespuestaRepository{DB: db}
}

func (r *RespuestaRepository) GetAll(page, pageSize int) ([]Respuesta, error) {
	var respuestas []Respuesta
	offset := (page - 1) * pageSize
	query := "SELECT * FROM respuestas LIMIT ? OFFSET ?"
	err := r.DB.Select(&respuestas, query, pageSize, offset)
	return respuestas, err
}

func (r *RespuestaRepository) GetAllByParticipanteIdAndNumeroAtencion(participanteId, numeroAtencion, page, pageSize int) ([]RespuestaDetalle, error) {
	var respuestas []RespuestaDetalle
	offset := (page - 1) * pageSize

	query := `
		SELECT 
			r.id_respuesta,
			r.id_participante,
			r.id_encuesta,
			r.id_pregunta,
			p.texto_pregunta,
			r.respuesta,
			r.created_at,
			r.updated_at,
			r.numero_atencion
		FROM respuestas r
		JOIN questions p ON r.id_pregunta = p.id 
		WHERE r.id_participante = ? AND r.numero_atencion = ?
		LIMIT ? OFFSET ?`

	err := r.DB.Select(&respuestas, query, participanteId, numeroAtencion, pageSize, offset)
	return respuestas, err
}

func (r *RespuestaRepository) GetByID(id int) (*Respuesta, error) {
	var respuesta Respuesta
	err := r.DB.Get(&respuesta, "SELECT * FROM respuestas WHERE id_respuesta = ?", id)
	return &respuesta, err
}

func (r *RespuestaRepository) Create(res *Respuesta) error {
	_, err := r.DB.NamedExec(`
		INSERT INTO respuestas (id_participante, id_encuesta, id_pregunta, respuesta, created_at, updated_at)
		VALUES (:id_participante, :id_encuesta, :id_pregunta, :respuesta, :created_at, :updated_at)`, res)
	return err
}

// func (r *RespuestaRepository) CreateBatch(respuestas []Respuesta) error {
// 	query := `
// 		INSERT INTO respuestas (id_participante, id_encuesta, id_pregunta, id_historial, respuesta, created_at, updated_at, numero_atencion)
// 		VALUES (:id_participante, :id_encuesta, :id_pregunta, :id_historial, :respuesta, :created_at, :updated_at, :numero_atencion)
// 	`
// 	_, err := r.DB.NamedExec(query, respuestas)
// 	return err
// }

func (r *RespuestaRepository) CreateBatch(respuestas []Respuesta) error {
	query := `
		INSERT INTO respuestas (
			id_participante, id_encuesta, id_pregunta, id_historial, 
			respuesta, created_at, updated_at, numero_atencion
		)
		VALUES (
			:id_participante, 
			CASE WHEN :id_encuesta IS NOT NULL THEN :id_encuesta ELSE NULL END, 
			:id_pregunta, :id_historial, 
			:respuesta, :created_at, :updated_at, :numero_atencion
		)
	`
	_, err := r.DB.NamedExec(query, respuestas)
	return err
}

func (r *RespuestaRepository) Update(id int, res *Respuesta) error {
	query := "UPDATE respuestas SET "
	params := make(map[string]interface{})

	if res.IDParticipante != 0 {
		query += "id_participante=:id_participante, "
		params["id_participante"] = res.IDParticipante
	}
	if res.IDEncuesta != 0 {
		query += "id_encuesta=:id_encuesta, "
		params["id_encuesta"] = res.IDEncuesta
	}
	if res.IDPregunta != 0 {
		query += "id_pregunta=:id_pregunta, "
		params["id_pregunta"] = res.IDPregunta
	}
	if res.Respuesta != "" {
		query += "respuesta=:respuesta, "
		params["respuesta"] = res.Respuesta
	}

	query = query[:len(query)-2] + " WHERE id_respuesta=:id_respuesta"
	params["id_respuesta"] = id

	_, err := r.DB.NamedExec(query, params)
	return err
}

func (r *RespuestaRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM respuestas WHERE id_respuesta = ?", id)
	return err
}

func (r *RespuestaRepository) GetResponsesPerParticipant(idParticipante int) ([]RespuestaDetalle, error) {
	query := `
		SELECT r.id_respuesta, r.id_participante, r.id_encuesta, r.id_pregunta, p.texto_pregunta, r.respuesta, r.created_at, r.updated_at FROM respuestas r JOIN questions p ON r.id_pregunta = id WHERE r.id_participante = ?;
	`
	var respuestas []RespuestaDetalle
	err := r.DB.Select(&respuestas, query, idParticipante)
	return respuestas, err
}
