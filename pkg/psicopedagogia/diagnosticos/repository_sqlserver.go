package diagnosticos

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DiagnosticoRepository struct {
	DB *sqlx.DB
}

func NewDiagnosticoRepository(db *sqlx.DB) *DiagnosticoRepository {
	return &DiagnosticoRepository{DB: db}
}

func (s *DiagnosticoRepository) GetAll() ([]*Diagnostico, error) {
	const sqlGetAll = `
		SELECT id, codigo, nombre, created_at, updated_at
		FROM diagnosticos
	`
	var diagnosticos []*Diagnostico
	err := s.DB.Select(&diagnosticos, sqlGetAll)
	if err != nil {
		return nil, err
	}
	return diagnosticos, nil
}

func (s *DiagnosticoRepository) GetByID(id int) (*Diagnostico, error) {
	const sqlGetByID = `
		SELECT id, codigo, nombre, created_at, updated_at
		FROM diagnosticos
		WHERE id = ?
	`

	d := Diagnostico{}
	err := s.DB.Get(&d, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

func (s *DiagnosticoRepository) Create(d *Diagnostico) (int64, error) {
	const sqlInsert = `
		INSERT INTO diagnosticos (codigo, nombre, created_at, updated_at)
		VALUES (?, ?, ?, ?)
	`

	now := time.Now()
	result, err := s.DB.Exec(sqlInsert, d.Codigo, d.Nombre, now, now)
	if err != nil {
		return 0, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (s *DiagnosticoRepository) Update(d *Diagnostico) error {
	const sqlUpdate = `
		UPDATE diagnosticos
		SET codigo = ?, nombre = ?, updated_at = ?
		WHERE id = ?
	`
	d.UpdatedAt = time.Now()
	_, err := s.DB.Exec(sqlUpdate, d.Codigo, d.Nombre, d.UpdatedAt, d.ID)
	return err
}

func (s *DiagnosticoRepository) Delete(id int) error {
	const sqlDelete = `
		DELETE FROM diagnosticos
		WHERE id = ?
	`
	_, err := s.DB.Exec(sqlDelete, id)
	return err
}
