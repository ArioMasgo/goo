package gpdfs

import (
	"github.com/jmoiron/sqlx"
)

type GpdfsRepository struct {
	DB *sqlx.DB
}

func NewGpdfsRepository(db *sqlx.DB) *GpdfsRepository {
	return &GpdfsRepository{DB: db}
}

func (r *GpdfsRepository) GetResponsesPerParticipant(idParticipante, numeroAtencion int) ([]DataPDFRespuesta, error) {
	query := `
        SELECT pr.texto_pregunta, r.respuesta
        FROM respuestas r
        JOIN questions pr ON r.id_pregunta = pr.id
        WHERE r.id_participante = ? AND r.numero_atencion = ?;
    `
	var respuestas []DataPDFRespuesta
	err := r.DB.Select(&respuestas, query, idParticipante, numeroAtencion)
	if err != nil {
		return nil, err
	}
	return respuestas, nil
}

func (r *GpdfsRepository) GetParticipantByID(idParticipante int) (*ParticipantePDF, error) {
	query := `
		SELECT * FROM participantes WHERE id_participante = ?;
	`
	var participante ParticipantePDF
	err := r.DB.Get(&participante, query, idParticipante)
	return &participante, err
}
