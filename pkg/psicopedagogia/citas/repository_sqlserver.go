package citas

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CitaRepository struct {
	DB *sqlx.DB
}

func NewCitaRepository(db *sqlx.DB) *CitaRepository {
	return &CitaRepository{DB: db}
}

func (r *CitaRepository) GetAll() ([]*Cita, error) {
	const sqlGetAll = `
		SELECT id, dni, nombre, apellido, facultad, fecha_inicio, fecha_fin, estado, created_at
		FROM citas
	`
	var citas []*Cita
	err := r.DB.Select(&citas, sqlGetAll)
	if err != nil {
		return nil, err
	}
	return citas, nil
}

func (r *CitaRepository) GetByID(id int) (*Cita, error) {
	const sqlGetByID = `
		SELECT id, dni, nombre, apellido, facultad, fecha_inicio, fecha_fin, estado, created_at
		FROM citas
		WHERE id = ?
	`

	c := Cita{}
	err := r.DB.Get(&c, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *CitaRepository) Create(c *Cita) (int64, error) {
	const sqlInsert = `
		INSERT INTO citas (dni, nombre, apellido, facultad, fecha_inicio, created_at, fecha_fin)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	now := time.Now()
	result, err := r.DB.Exec(sqlInsert, c.DNI, c.Nombre, c.Apellido, c.Facultad, c.FechaInicio, now, c.FechaFin)
	if err != nil {
		return 0, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (r *CitaRepository) Update(c *Cita) error {
	const sqlUpdate = `
		UPDATE citas
		SET dni = ?, nombre = ?, apellido = ?, facultad = ?, fecha_inicio = ?, fecha_fin = ?
		WHERE id = ?
	`

	_, err := r.DB.Exec(sqlUpdate, c.DNI, c.Nombre, c.Apellido, c.Facultad, c.FechaInicio, c.FechaFin, c.ID)
	return err
}

func (r *CitaRepository) Delete(id int) error {
	const sqlDelete = `
		DELETE FROM citas
		WHERE id = ?
	`
	_, err := r.DB.Exec(sqlDelete, id)
	return err
}

func (r *CitaRepository) ExisteCitaEnFecha(fecha time.Time, dni string) (bool, error) {
	const sqlCheckFecha = `
		SELECT COUNT(*) 
		FROM citas 
		WHERE dni = ? 
		AND estado = 'Pendiente'
		AND fecha_inicio <= ? 
		AND fecha_fin >= ?
	`

	var count int
	err := r.DB.Get(&count, sqlCheckFecha, dni, fecha, fecha)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
