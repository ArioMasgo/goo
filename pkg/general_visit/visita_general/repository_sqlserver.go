package visita_general

import (
	"database/sql"
	"dbu-api/internal/logger"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type sqlserver struct {
	DB   *sqlx.DB
	TxID string
}

func newVisitaGeneralSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s *sqlserver) create(m *VisitaGeneral) error {
	const sqlInsert = `INSERT INTO visitaGeneral (id, tipo_usuario, codigo_estudiante, dni, nombre_completo, genero, edad, escuela, area, motivo_atencion, descripcion_motivo, url_imagen, departamento, provincia, distrito, lugar_atencion, created_by, created_at, updated_by, updated_at) 
              VALUES (:id, :tipo_usuario, :codigo_estudiante, :dni, :nombre_completo, :genero, :edad, :escuela, :area, :motivo_atencion, :descripcion_motivo, :url_imagen, :departamento, :provincia, :distrito, :lugar_atencion, :created_by, :created_at, :updated_by, :updated_at)`
	rs, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s *sqlserver) getAll() ([]*VisitaGeneral, error) {
	const sqlGetAll = `SELECT id, tipo_usuario, codigo_estudiante, dni, nombre_completo, genero, edad, escuela, area, motivo_atencion, descripcion_motivo, url_imagen, departamento, provincia, distrito, lugar_atencion, created_by, created_at, updated_by, updated_at FROM visitaGeneral`
	var ms []*VisitaGeneral

	if err := s.DB.Ping(); err != nil {
		logger.Error.Printf("%s - Database connection error: %v", s.TxID, err)
		return nil, fmt.Errorf("database connection error: %v", err)
	}
	
	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		logger.Error.Printf("%s - Error querying visitaGeneral table: %v", s.TxID, err)
		if err == sql.ErrNoRows {
			return []*VisitaGeneral{}, nil 
		}
		if strings.Contains(err.Error(), "bad connection") || strings.Contains(err.Error(), "connection refused") {
			logger.Error.Printf("%s - Database connection is bad or refused: %v", s.TxID, err)
			return nil, fmt.Errorf("database connection error: %v", err)
		}
		
		var exists bool
		checkTableSQL := "SHOW TABLES LIKE 'visitaGeneral'"
		tableErr := s.DB.Get(&exists, checkTableSQL)
		if tableErr != nil {
			logger.Error.Printf("%s - Error checking if table exists: %v", s.TxID, tableErr)
			
			if strings.Contains(tableErr.Error(), "bad connection") || strings.Contains(tableErr.Error(), "connection refused") {
				return nil, fmt.Errorf("database connection error: %v", tableErr)
			}
			
			return nil, fmt.Errorf("error checking database table: %v", tableErr)
		}
		
		if !exists {
			logger.Error.Printf("%s - Table 'visitaGeneral' does not exist", s.TxID)
			return nil, fmt.Errorf("table 'visitaGeneral' does not exist")
		}
		
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	if len(ms) == 0 {
		logger.Info.Printf("%s - No records found in visitaGeneral table", s.TxID)
		return []*VisitaGeneral{}, nil 
	}
	
	return ms, nil
}

func (s *sqlserver) update(m *VisitaGeneral) error {
	const sqlUpdate = `UPDATE visitaGeneral SET tipo_usuario = :tipo_usuario, codigo_estudiante = :codigo_estudiante, 
		dni = :dni, nombre_completo = :nombre_completo, genero = :genero, edad = :edad, 
		escuela = :escuela, area = :area, motivo_atencion = :motivo_atencion, 
		descripcion_motivo = :descripcion_motivo, url_imagen = :url_imagen, departamento = :departamento, 
		provincia = :provincia, distrito = :distrito, lugar_atencion = :lugar_atencion, 
		updated_by = :updated_by, updated_at = :updated_at 
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

func (s *sqlserver) getByID(id string) (*VisitaGeneral, error) {
	const sqlGetByID = `SELECT id, tipo_usuario, codigo_estudiante, dni, nombre_completo, genero, edad, escuela, area, motivo_atencion, descripcion_motivo, url_imagen, departamento, provincia, distrito, lugar_atencion, created_by, created_at, updated_by, updated_at FROM visitaGeneral WHERE id = ?`
	var m VisitaGeneral
	if err := s.DB.Ping(); err != nil {
		logger.Error.Printf("%s - Database connection error: %v", s.TxID, err)
		return nil, fmt.Errorf("database connection error: %v", err)
	}

	err := s.DB.Get(&m, sqlGetByID, id)
	if err != nil {
		logger.Error.Printf("%s - Error querying visitaGeneral table: %v", s.TxID, err)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("record not found")
		}

		if strings.Contains(err.Error(), "bad connection") || strings.Contains(err.Error(), "connection refused") {
			logger.Error.Printf("%s - Database connection is bad or refused: %v", s.TxID, err)
			return nil, fmt.Errorf("database connection error: %v", err)
		}
		
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	
	return &m, nil
}

func (s *sqlserver) delete(id string) error {
	const sqlDelete = `DELETE FROM visitaGeneral WHERE id = :id`
	m := VisitaGeneral{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s *sqlserver) getAllDepartments() ([]*Departamento, error) {
	const sqlGetDepartments = `SELECT id, name, created_at, updated_at FROM departaments ORDER BY name`
	var departments []*Departamento
	
	err := s.DB.Select(&departments, sqlGetDepartments)
	if err != nil {
		logger.Error.Printf("%s - Error querying departments: %v", s.TxID, err)
		return nil, fmt.Errorf("error querying departments: %v", err)
	}
	
	return departments, nil
}
func (s *sqlserver) getProvincesByDepartment(departmentID string) ([]*Provincia, error) {
	const sqlGetProvinces = `SELECT id, name, departament_id, created_at, updated_at FROM provinces WHERE departament_id = ? ORDER BY name`
	var provinces []*Provincia
	
	err := s.DB.Select(&provinces, sqlGetProvinces, departmentID)
	if err != nil {
		logger.Error.Printf("%s - Error querying provinces for department %s: %v", s.TxID, departmentID, err)
		return nil, fmt.Errorf("error querying provinces: %v", err)
	}
	
	return provinces, nil
}
func (s *sqlserver) getDistrictsByProvince(provinceID string) ([]*Distrito, error) {
	const sqlGetDistricts = `SELECT id, name, province_id, created_at, updated_at FROM districts WHERE province_id = ? ORDER BY name`
	var districts []*Distrito
	
	err := s.DB.Select(&districts, sqlGetDistricts, provinceID)
	if err != nil {
		logger.Error.Printf("%s - Error querying districts for province %s: %v", s.TxID, provinceID, err)
		return nil, fmt.Errorf("error querying districts: %v", err)
	}
	
	return districts, nil
}

func (s *sqlserver) getLocationHierarchy() (*LocationResponse, error) {
	departments, err := s.getAllDepartments()
	if err != nil {
		return nil, err
	}
	
	var response LocationResponse
	response.Departamentos = make([]DepartamentoWithProvincias, 0, len(departments))
	
	for _, dept := range departments {
		deptWithProvinces := DepartamentoWithProvincias{
			ID:         dept.ID,
			Name:       dept.Name,
			Provincias: make([]ProvinciaWithDistritos, 0),
		}
		provinces, err := s.getProvincesByDepartment(dept.ID)
		if err != nil {
			logger.Warning.Printf("%s - Error getting provinces for department %s: %v", s.TxID, dept.ID, err)
			continue
		}
		
		for _, prov := range provinces {
			provWithDistricts := ProvinciaWithDistritos{
				ID:        prov.ID,
				Name:      prov.Name,
				Distritos: make([]Distrito, 0),
			}
			
			districts, err := s.getDistrictsByProvince(prov.ID)
			if err != nil {
				logger.Warning.Printf("%s - Error getting districts for province %s: %v", s.TxID, prov.ID, err)
				continue
			}
			
			for _, dist := range districts {
				provWithDistricts.Distritos = append(provWithDistricts.Distritos, Distrito{
					ID:         dist.ID,
					Name:       dist.Name,
					ProvinceID: dist.ProvinceID,
					CreatedAt:  dist.CreatedAt,
					UpdatedAt:  dist.UpdatedAt,
				})
			}
			
			deptWithProvinces.Provincias = append(deptWithProvinces.Provincias, provWithDistricts)
		}
		
		response.Departamentos = append(response.Departamentos, deptWithProvinces)
	}
	
	return &response, nil
}