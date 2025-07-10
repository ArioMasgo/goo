package encuestas

import (
	"github.com/jmoiron/sqlx"
)

type EncuestaRepository struct {
	DB *sqlx.DB
}

func NewEncuestaRepository(db *sqlx.DB) *EncuestaRepository {
	return &EncuestaRepository{DB: db}
}

func (r *EncuestaRepository) GetAll(page, pageSize int) ([]Encuesta, error) {
	var encuestas []Encuesta
	offset := (page - 1) * pageSize
	query := "SELECT * FROM encuestas LIMIT ? OFFSET ?"
	err := r.DB.Select(&encuestas, query, pageSize, offset)
	return encuestas, err
}

func (r *EncuestaRepository) GetByID(id int) (*Encuesta, error) {
	var encuesta Encuesta
	err := r.DB.Get(&encuesta, "SELECT * FROM encuestas WHERE id_encuesta = ?", id)
	return &encuesta, err
}

func (r *EncuestaRepository) Create(e *Encuesta) (int64, error) {
	query := `
		INSERT INTO encuestas (nombre_encuesta, descripcion, estado, fecha_inicio, fecha_fin)
		VALUES (:nombre_encuesta, :descripcion, :estado, :fecha_inicio, :fecha_fin)`

	result, err := r.DB.NamedExec(query, e)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *EncuestaRepository) Update(id int, e *Encuesta) error {
	query := "UPDATE encuestas SET "
	params := make(map[string]interface{})

	if e.Nombre != "" {
		query += "nombre_encuesta=:nombre_encuesta, "
		params["nombre_encuesta"] = e.Nombre
	}
	if e.Descripcion != nil {
		query += "descripcion=:descripcion, "
		params["descripcion"] = *e.Descripcion
	}
	if e.Estado != "" {
		query += "estado=:estado, "
		params["estado"] = e.Estado
	}
	if e.FechaInicio != nil {
		query += "fecha_inicio=:fecha_inicio, "
		params["fecha_inicio"] = *e.FechaInicio
	}
	if e.FechaFin != nil {
		query += "fecha_fin=:fecha_fin, "
		params["fecha_fin"] = *e.FechaFin
	}

	query = query[:len(query)-2] + " WHERE id_encuesta=:id_encuesta"
	params["id_encuesta"] = id

	_, err := r.DB.NamedExec(query, params)
	return err
}

func (r *EncuestaRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM encuestas WHERE id_encuesta = ?", id)
	return err
}

func (r *EncuestaRepository) GetActiveSRQ() (int, string, bool, error) {
	var encuesta struct {
		ID       int    `db:"id_encuesta"`
		FechaFin string `db:"fecha_fin"`
	}

	query := `SELECT id_encuesta, fecha_fin FROM encuestas 
			  WHERE estado = 'activa' 
			  AND nombre_encuesta = 'S.R.Q' 
			  AND fecha_fin >= CURDATE() 
			  LIMIT 1`

	err := r.DB.Get(&encuesta, query)
	if err != nil {
		return 0, "", false, err
	}

	return encuesta.ID, encuesta.FechaFin, true, nil
}
