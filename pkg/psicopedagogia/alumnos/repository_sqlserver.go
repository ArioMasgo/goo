package alumnos

import (
	"database/sql"
	"dbu-api/pkg/psicopedagogia/participantes"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type sqlserver struct {
	DB   *sqlx.DB
	TxID string
}

func newEstudianteSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s *sqlserver) getByDni(dni string) (*Estudiante, error) {
	const sqlGetByDni = `
		SELECT 
			id, 
			codigo_estudiante, 
			DNI, 
			nombres, 
			apellido_paterno, 
			apellido_materno, 
			sexo, 
			facultad, 
			escuela_profesional, 
			ultimo_semestre, 
			modalidad_ingreso, 
			lugar_procedencia, 
			lugar_nacimiento, 
			edad, 
			correo_institucional, 
			direccion, 
			fecha_nacimiento, 
			correo_personal, 
			celular_estudiante, 
			celular_padre, 
			estado_matricula, 
			creditos_matriculados, 
			num_semestres_cursados
		FROM alumnos 
		WHERE DNI = ?
	`

	e := Estudiante{}
	err := s.DB.Get(&e, sqlGetByDni, dni)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &e, nil
}

func (s *sqlserver) getNameByDni(dni string) (*BasicEstudiante, error) {
	const sqlGetByDni = `
		SELECT 
			id, 
			nombres, 
			apellido_paterno, 
			apellido_materno
		FROM alumnos 
		WHERE DNI = ?
	`

	e := BasicEstudiante{}
	err := s.DB.Get(&e, sqlGetByDni, dni)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &e, nil
}

func (r *sqlserver) GetParticipanteByDNI(dni string, tipoParticipante string) (*participantes.Participante, error) {
	var participante participantes.Participante
	err := r.DB.Get(&participante, "SELECT * FROM participantes WHERE dni = ? AND tipo_participante = ?", dni, tipoParticipante)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &participante, nil
}
