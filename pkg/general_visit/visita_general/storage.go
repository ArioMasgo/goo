package visita_general

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServiceVisitaGeneralRepository interface {
	create(m *VisitaGeneral) error
	getAll() ([]*VisitaGeneral, error)
	getByID(id string) (*VisitaGeneral, error)
	update(m *VisitaGeneral) error
	delete(id string) error
	
	// Métodos para ubicaciones
	getAllDepartments() ([]*Departamento, error)
	getProvincesByDepartment(departmentID string) ([]*Provincia, error)
	getDistrictsByProvince(provinceID string) ([]*Distrito, error)
	getLocationHierarchy() (*LocationResponse, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServiceVisitaGeneralRepository {
	return newVisitaGeneralSqlServerRepository(db, txID)
}