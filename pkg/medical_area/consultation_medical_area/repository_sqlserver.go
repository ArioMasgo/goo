package consultation_medical_area

import (
	"database/sql"
	"dbu-api/internal/models"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type sqlserver struct {
	DB   *sqlx.DB
	TxID string
}

func newConsultationMedicalAreaSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *ConsultationMedicalArea) error {
	const sqlInsertConsulta = `INSERT INTO consultas_areas_medicas (id, paciente_id, fecha_consulta, area_medica, user_creator, created_at, updated_at) VALUES (:id, :paciente_id, :fecha_consulta, :area_medica, :user_creator, :created_at, :updated_at)`

	rs, err := s.DB.NamedExec(sqlInsertConsulta, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *ConsultationMedicalArea) error {
	const sqlUpdate = `UPDATE consultas_areas_medicas SET paciente_id = :paciente_id, fecha_consulta = :fecha_consulta, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id`
	_, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	return nil
}

func (s sqlserver) delete(id string) error {
	const psqlDelete = `DELETE FROM consultas_areas_medicas WHERE id = :id`
	m := ConsultationMedicalArea{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*ConsultationMedicalArea, error) {
	const sqlGetByID = `SELECT id, paciente_id, fecha_consulta, area_medica, deleted_at, user_creator, created_at, updated_at
                        FROM consultas_areas_medicas WHERE id = ?`

	mdl := ConsultationMedicalArea{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &mdl, nil
}

func (s sqlserver) getAll() ([]*ConsultationMedicalArea, error) {
	var ms []*ConsultationMedicalArea
	const sqlGetAll = `SELECT id, paciente_id, fecha_consulta, area_medica, created_at, updated_at FROM consultas_areas_medicas`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}

func (s sqlserver) GetAllByPatientID(id string) ([]*ConsultationMedicalArea, error) {
	var ms []*ConsultationMedicalArea
	const sqlGetAll = `SELECT id, paciente_id, fecha_consulta, area_medica, created_at, updated_at FROM consultas_areas_medicas WHERE paciente_id = ?`

	err := s.DB.Select(&ms, sqlGetAll, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}

func (s sqlserver) getByPatientDNI(dni string) ([]*ConsultationMedicalArea, error) {
	var ms []*ConsultationMedicalArea
	const sqlGetAll = `SELECT cam.id, cam.paciente_id, cam.fecha_consulta, cam.area_medica, cam.created_at, cam.updated_at FROM consultas_areas_medicas cam JOIN pacientes p ON p.id = cam.paciente_id WHERE p.dni = ?`

	err := s.DB.Select(&ms, sqlGetAll, dni)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}

func (s sqlserver) getAllByNursingDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error) {
	var ms []*models.ConsultationPatientsMedicalAreaExcel
	const sqlGetAll = `
		SELECT 
			cam.id as id,
			cam.fecha_consulta as fecha_consulta,
			p.tipo_persona as tipo_persona, 
			coalesce(p.codigo_sga,'') as codigo_sga,
			p.dni as dni, 
			CONCAT(p.nombres, ' ', p.apellidos) AS nombre_completo, 
			p.sexo as sexo,
			p.fecha_nacimiento as fecha_nacimiento,
			p.numero_celular as numero_celular,
			TRIM(BOTH ', ' FROM CONCAT_WS(', ',
				IF(epr.id IS NOT NULL, 'Procedimiento', NULL),
				IF(ees.id IS NOT NULL, 'Examen sexualidad', NULL),
				IF(eel.id IS NOT NULL, 'Examen laboratorio', NULL),
				IF(eev.id IS NOT NULL, 'Examen visual', NULL),
				IF(etm.id IS NOT NULL, 'Tratamiento medicamentoso', NULL),
				IF(eef.id IS NOT NULL, 'Examen físico', NULL),
				IF(ev.id IS NOT NULL, 'Vacuna', NULL),     
				IF(mai.id IS NOT NULL, 'Consulta enfermería', NULL)
			)) AS servicios,
			coalesce(p.procedencia,'') as procedencia,
			coalesce(p.direccion) as direccion_residencia,
			coalesce(epr.procedimiento,'') as procedimiento,
			coalesce(epr.numero_recibo,'') as numero_recibo,
			coalesce(epr.costo,'') as costo
		FROM consultas_areas_medicas cam
		JOIN pacientes p ON p.id = cam.paciente_id
		LEFT JOIN enfermeria_procedimientos_realizados epr ON epr.consulta_enfermeria_id = cam.id
		LEFT JOIN enfermeria_examen_sexualidad ees ON ees.consulta_enfermeria_id = cam.id
		LEFT JOIN enfermeria_examen_laboratorio eel ON eel.consulta_enfermeria_id = cam.id
		LEFT JOIN enfermeria_examen_visual eev ON eev.consulta_enfermeria_id = cam.id
		LEFT JOIN enfermeria_tratamiento_medicamentoso etm ON etm.consulta_enfermeria_id = cam.id
		LEFT JOIN enfermeria_examen_fisico eef ON eef.consulta_enfermeria_id = cam.id
		LEFT JOIN enfermeria_vacuna ev ON ev.consulta_enfermeria_id = cam.id
		LEFT JOIN medicina_atencion_integral mai ON mai.consulta_id = cam.id
		WHERE cam.area_medica = ?
		AND cam.created_at >= ?
		AND cam.created_at <= ?
	`

	err := s.DB.Select(&ms, sqlGetAll, area_medica, fecha_inicio, fecha_fin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}

func (s sqlserver) getAllByDentistryDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error) {
	var ms []*models.ConsultationPatientsMedicalAreaExcel
	const sqlGetAll = `
			SELECT 
				cam.id as id,
				cam.fecha_consulta as fecha_consulta,
				p.tipo_persona as tipo_persona,
				p.dni as dni, 
				CONCAT(p.nombres, ' ', p.apellidos) AS nombre_completo, 
				p.sexo as sexo,
				p.fecha_nacimiento as fecha_nacimiento,
				coalesce(p.ocupacion,'') as ocupacion,
				p.numero_celular as numero_celular,
				TRIM(BOTH ', ' FROM CONCAT_WS(', ',
					IF(oc.id IS NOT NULL, 'Consulta odontología', NULL),     
					IF(oe.id IS NOT NULL, 'Examen', NULL),
					IF(op.id IS NOT NULL, 'Procedimiento', NULL)
				)) AS servicios,
				coalesce(op.tipo_procedimiento,'') as tipo_procedimiento,
				coalesce(op.recibo,'') as recibo,
				coalesce(op.costo,'') as costo,
				coalesce(op.fecha_pago,'') as fecha_pago,
				coalesce(op.pieza_dental,'') as pieza_dental
			FROM consultas_areas_medicas cam
			JOIN pacientes p ON p.id = cam.paciente_id
			LEFT JOIN odontologia_consulta oc ON oc.consulta_odontologia_id = cam.id 
			LEFT JOIN odontologia_examen oe ON oe.consulta_odontologia_id = cam.id 
			LEFT JOIN odontologia_procedimiento op ON op.consulta_odontologia_id = cam.id
			WHERE cam.area_medica = ?
			AND cam.created_at >= ?
			AND cam.created_at <= ?
	`

	err := s.DB.Select(&ms, sqlGetAll, area_medica, fecha_inicio, fecha_fin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}

func (s sqlserver) getAllByMedicalDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error) {
	var ms []*models.ConsultationPatientsMedicalAreaExcel
	const sqlGetAll = `
		SELECT 
				cam.id as id,
				cam.fecha_consulta as fecha_consulta,
				p.tipo_persona as tipo_persona,
				p.dni as dni,
				CONCAT(p.nombres, ' ', p.apellidos) AS nombre_completo, 
				p.sexo as sexo,
				p.fecha_nacimiento as fecha_nacimiento,
				p.numero_celular as numero_celular,
				TRIM(BOTH ', ' FROM CONCAT_WS(', ',
					IF(mcg.id IS NOT NULL, 'Consulta medicina', NULL)
				)) AS servicios
			FROM consultas_areas_medicas cam
			JOIN pacientes p ON p.id = cam.paciente_id
			LEFT JOIN medicina_consulta_general mcg ON mcg.consulta_id = cam.id	
			WHERE cam.area_medica = ?
			AND cam.created_at >= ?
			AND cam.created_at <= ?
	`

	err := s.DB.Select(&ms, sqlGetAll, area_medica, fecha_inicio, fecha_fin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
