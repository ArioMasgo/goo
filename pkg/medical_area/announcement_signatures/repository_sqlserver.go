package announcement_signatures

import (
	"database/sql"
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

func newAnnouncementSignaturesSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s *sqlserver) create(m *AnnouncementSignatures) error {
	const sqlInsert = `INSERT INTO firmas_convocatoria_area_medica (id, convocatoria_id, paciente_id, firma_enfermeria, firma_medicina, firma_odontologia, firma_psicologia, user_creator, created_at, updated_at) VALUES (:id, :convocatoria_id, :paciente_id, :firma_enfermeria, :firma_medicina, :firma_odontologia, :firma_psicologia, :user_creator, :created_at, :updated_at)`
	rs, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s *sqlserver) update(m *AnnouncementSignatures) error {
	const sqlUpdate = `UPDATE firmas_convocatoria_area_medica SET convocatoria_id = :convocatoria_id, paciente_id = :paciente_id, firma_enfermeria = :firma_enfermeria, firma_medicina = :firma_medicina, firma_odontologia = :firma_odontologia, firma_psicologia = :firma_psicologia, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id`
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s *sqlserver) delete(id string) error {
	const psqlDelete = `DELETE FROM firmas_convocatoria_area_medica WHERE id = :id`
	m := AnnouncementSignatures{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s *sqlserver) getByID(id string) (*AnnouncementSignatures, error) {
	const sqlGetByID = `SELECT id, convocatoria_id, paciente_id, firma_enfermeria, firma_medicina, firma_odontologia, firma_psicologia, is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM firmas_convocatoria_area_medica WHERE id = ?`
	mdl := AnnouncementSignatures{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s *sqlserver) getByIDPatient(id string) (*AnnouncementSignatures, error) {
	const sqlGetByIDPatient = `SELECT id, convocatoria_id, paciente_id, firma_enfermeria, firma_medicina, firma_odontologia, firma_psicologia, is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM firmas_convocatoria_area_medica WHERE paciente_id = ?`
	mdl := AnnouncementSignatures{}
	err := s.DB.Get(&mdl, sqlGetByIDPatient, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s *sqlserver) getByDNIPatient(id string) (*AnnouncementSignatures, error) {
	const sqlGetByIDPatient = `SELECT fcam.id, fcam.convocatoria_id, fcam.paciente_id, fcam.firma_enfermeria, fcam.firma_medicina, fcam.firma_odontologia, fcam.firma_psicologia, fcam.is_deleted, fcam.user_deleted, fcam.deleted_at, fcam.user_creator, fcam.created_at, fcam.updated_at FROM firmas_convocatoria_area_medica fcam JOIN pacientes p on p.id = fcam.paciente_id where p.dni = ?`
	mdl := AnnouncementSignatures{}
	err := s.DB.Get(&mdl, sqlGetByIDPatient, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s *sqlserver) getAll() ([]*AnnouncementSignatures, error) {
	var ms []*AnnouncementSignatures
	const sqlGetAll = `SELECT id, convocatoria_id, paciente_id, firma_enfermeria, firma_medicina, firma_odontologia, firma_psicologia, is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM firmas_convocatoria_area_medica`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) GetAnnouncement() (*Announcement, error) {
	const sqlGetAnnouncement = `SELECT id, user_id, fecha_inicio, fecha_fin, nombre FROM convocatorias order by created_at desc limit 1`
	mdl := Announcement{}
	err := s.DB.Get(&mdl, sqlGetAnnouncement)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}
