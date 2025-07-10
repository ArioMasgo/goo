package patients

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	TxID string
}

func newPatientsSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s *sqlserver) existsByDNI(dni string) (bool, error) {
	var exists bool = false
	var count int
	const sqlExistsByDNI = `SELECT COUNT(1) FROM pacientes WHERE dni = ?`
	err := s.DB.QueryRow(sqlExistsByDNI, dni).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		exists = true
	}
	return exists, nil
}

func (s *sqlserver) create(m *Patients) error {
	const sqlInsert = `INSERT INTO pacientes (id, codigo_sga, dni, nombres, apellidos, sexo, edad, estado_civil,grupo_sanguineo, fecha_nacimiento, lugar_nacimiento, procedencia,escuela_profesional, ocupacion, correo_electronico,numero_celular, direccion, tipo_persona, factor_rh, alergias, ram, user_creator, created_at, updated_at) 
        VALUES (:id,:codigo_sga,:dni,:nombres,:apellidos,:sexo,:edad,:estado_civil,:grupo_sanguineo,:fecha_nacimiento,:lugar_nacimiento,:procedencia,:escuela_profesional,:ocupacion,:correo_electronico,:numero_celular,:direccion,:tipo_persona,:factor_rh,:alergias,:ram,:user_creator,:created_at,:updated_at)`
	rs, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s *sqlserver) update(m *Patients) error {
	const sqlUpdate = `UPDATE pacientes SET 
        codigo_sga = :codigo_sga, 
        dni = :dni, 
        nombres = :nombres, 
        apellidos = :apellidos, 
        sexo = :sexo, 
        edad = :edad, 
        estado_civil = :estado_civil, 
        grupo_sanguineo = :grupo_sanguineo, 
        fecha_nacimiento = :fecha_nacimiento, 
        lugar_nacimiento = :lugar_nacimiento, 
        procedencia = :procedencia, 
        escuela_profesional = :escuela_profesional, 
        ocupacion = :ocupacion, 
        correo_electronico = :correo_electronico, 
        numero_celular = :numero_celular, 
        direccion = :direccion, 
        tipo_persona = :tipo_persona, 
        factor_rh = :factor_rh, 
        alergias = :alergias, 
        ram = :ram,
        updated_at = now() 
    WHERE id = :id`
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
	const psqlDelete = `DELETE FROM pacientes WHERE id = :id`
	m := Patients{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s *sqlserver) getByID(id string) (*Patients, error) {
	const sqlGetByID = `SELECT id , codigo_sga, dni, nombres, apellidos, sexo, edad, estado_civil, grupo_sanguineo, fecha_nacimiento, lugar_nacimiento, procedencia, escuela_profesional, ocupacion, correo_electronico, numero_celular, direccion, tipo_persona, factor_rh, alergias, ram, created_at, updated_at FROM pacientes WHERE id = ?`
	mdl := Patients{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s *sqlserver) getByDNI(dni string) (*Patients, error) {
	const sqlGetByID = `SELECT id , codigo_sga, dni, nombres, apellidos, sexo, edad, estado_civil, grupo_sanguineo, fecha_nacimiento, lugar_nacimiento, procedencia, escuela_profesional, ocupacion, correo_electronico, numero_celular, direccion, tipo_persona, factor_rh, alergias, ram, created_at, updated_at FROM pacientes WHERE dni = ?`
	mdl := Patients{}
	err := s.DB.Get(&mdl, sqlGetByID, dni)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s *sqlserver) getAll() ([]*Patients, error) {
	var ms []*Patients
	const sqlGetAll = `SELECT id , codigo_sga, dni, nombres, apellidos, sexo, edad, estado_civil, grupo_sanguineo, fecha_nacimiento, lugar_nacimiento, procedencia, escuela_profesional, ocupacion, correo_electronico, numero_celular, direccion, tipo_persona, factor_rh, alergias, ram, created_at, updated_at FROM pacientes`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) countPaginationPatients(dni, names, surnames string) (int64, error) {
	var total int64

	query := `SELECT COUNT(*)
		FROM pacientes`

	query, params := prepareSearchPatientQuery(query, dni, names, surnames)

	err := s.DB.Get(&total, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}
	return total, nil
}

func (s *sqlserver) searchPaginationPatients(dni, names, surnames string, limit, offset int64) ([]*Patients, error) {
	var ms []*Patients

	query := `SELECT id, codigo_sga, dni, nombres, apellidos, sexo, edad, estado_civil, 
		grupo_sanguineo, fecha_nacimiento, lugar_nacimiento, procedencia, escuela_profesional, 
		ocupacion, correo_electronico, numero_celular, direccion, tipo_persona, factor_rh, 
		alergias, ram, created_at, updated_at 
		FROM pacientes`

	query, params := prepareSearchPatientQuery(query, dni, names, surnames)

	query += " LIMIT ? OFFSET ?"

	params = append(params, limit, offset)

	err := s.DB.Select(&ms, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return ms, nil

}

func prepareSearchPatientQuery(query, dni, names, surnames string) (string, []interface{}) {
	var conditions []string
	var params []interface{}

	if dni != "" {
		conditions = append(conditions, "dni LIKE ?")
		params = append(params, "%"+dni+"%")
	}
	if names != "" {
		conditions = append(conditions, "nombres LIKE ?")
		params = append(params, "%"+names+"%")
	}
	if surnames != "" {
		conditions = append(conditions, "apellidos LIKE ?")
		params = append(params, "%"+surnames+"%")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	return query, params
}
