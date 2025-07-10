package consultation_integral_attention

import (
	"database/sql"
	"dbu-api/internal/models"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	TxID string
}

func newConsultationIntegralAttentionSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *ConsultationIntegralAttention) error {
	const sqlInsertRevision = `INSERT INTO medicina_atencion_integral (id, consulta_id, fecha, hora, edad, motivo_consulta, tiempo_enfermedad, apetito, sed, suenio, estado_animo, orina, deposiciones, temperatura, presion_arterial, frecuencia_cardiaca, frecuencia_respiratoria, peso, talla, indice_masa_corporal, diagnostico, tratamiento, examenes_axuliares, referencia, observacion, numero_recibo, costo, fecha_pago, is_deleted, user_creator, created_at, updated_at) VALUES (:id, :consulta_id, :fecha, :hora, :edad, :motivo_consulta, :tiempo_enfermedad, :apetito, :sed, :suenio, :estado_animo, :orina, :deposiciones, :temperatura, :presion_arterial, :frecuencia_cardiaca, :frecuencia_respiratoria, :peso, :talla, :indice_masa_corporal, :diagnostico, :tratamiento, :examenes_axuliares, :referencia, :observacion, :numero_recibo, :costo, :fecha_pago, :is_deleted, :user_creator, :created_at, :updated_at)`
	rs, err := s.DB.NamedExec(sqlInsertRevision, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}
func (s sqlserver) update(m *ConsultationIntegralAttention) error {
	const sqlUpdate = `UPDATE medicina_atencion_integral SET fecha = :fecha, hora = :hora, edad = :edad, motivo_consulta = :motivo_consulta, tiempo_enfermedad = :tiempo_enfermedad, apetito = :apetito, sed = :sed, suenio = :suenio, estado_animo = :estado_animo, orina = :orina, deposiciones = :deposiciones, temperatura = :temperatura, presion_arterial = :presion_arterial, frecuencia_cardiaca = :frecuencia_cardiaca, frecuencia_respiratoria = :frecuencia_respiratoria, peso = :peso, talla = :talla, indice_masa_corporal = :indice_masa_corporal, diagnostico = :diagnostico, tratamiento = :tratamiento, examenes_axuliares = :examenes_axuliares, referencia = :referencia, observacion = :observacion, numero_recibo = :numero_recibo, costo = :costo, fecha_pago = :fecha_pago, updated_at = :updated_at WHERE id = :id`
	_, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	return nil
}

func (s sqlserver) delete(id string) error {
	const psqlDelete = `DELETE FROM medicina_atencion_integral WHERE id = :id`
	m := ConsultationIntegralAttention{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) deleteByIDConsultation(consulta_id string) error {
	const psqlDelete = `DELETE FROM medicina_atencion_integral WHERE consulta_id = :consulta_id`
	m := ConsultationIntegralAttention{IDConsulta: consulta_id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*ConsultationIntegralAttention, error) {
	const sqlGetByID = `SELECT id , consulta_id, fecha, hora, edad, motivo_consulta, tiempo_enfermedad, apetito, sed, suenio, estado_animo, orina, deposiciones, temperatura, presion_arterial, frecuencia_cardiaca, frecuencia_respiratoria, peso, talla, indice_masa_corporal, diagnostico, tratamiento, examenes_axuliares, referencia, observacion, numero_recibo, costo, fecha_pago, created_at, updated_at FROM medicina_atencion_integral WHERE id = ?`
	mdl := ConsultationIntegralAttention{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getByIDConsultation(id string) (*ConsultationIntegralAttention, error) {
	const sqlGetByID = `SELECT id , id, consulta_id, fecha, hora, edad, motivo_consulta, tiempo_enfermedad, apetito, sed, suenio, estado_animo, orina, deposiciones, temperatura, presion_arterial, frecuencia_cardiaca, frecuencia_respiratoria, peso, talla, indice_masa_corporal, diagnostico, tratamiento, examenes_axuliares, referencia, observacion, numero_recibo, costo, fecha_pago, created_at, updated_at FROM medicina_atencion_integral WHERE consulta_id = ?`
	mdl := ConsultationIntegralAttention{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getAll() ([]*ConsultationIntegralAttention, error) {
	var ms []*ConsultationIntegralAttention
	const sqlGetAll = `SELECT id, consulta_id, fecha, hora, edad, motivo_consulta, tiempo_enfermedad, apetito, sed, suenio, estado_animo, orina, deposiciones, temperatura, presion_arterial, frecuencia_cardiaca, frecuencia_respiratoria, peso, talla, indice_masa_corporal, diagnostico, tratamiento, examenes_axuliares, referencia, observacion, numero_recibo, costo, fecha_pago, created_at, updated_at FROM medicina_atencion_integral`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s sqlserver) getAllByDateExcel(fecha_inicio, fecha_fin string) ([]*models.ConsultationIntegralAttentionExcel, error) {
	var ms []*models.ConsultationIntegralAttentionExcel
	const sqlGetAll = `
		SELECT 
		cam.id as id, 
		cam.fecha_consulta as fecha_consulta, 
		p.tipo_persona as tipo_persona,
		p.escuela_profesional as escuela_profesional,
		p.sexo as sexo
	FROM consultas_areas_medicas cam
	JOIN pacientes p ON p.id = cam.paciente_id
	JOIN medicina_atencion_integral mai ON mai.consulta_id = cam.id
	AND cam.created_at >= ?
	AND cam.created_at <= ?
	`

	err := s.DB.Select(&ms, sqlGetAll, fecha_inicio, fecha_fin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
