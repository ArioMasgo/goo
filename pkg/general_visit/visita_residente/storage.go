package visita_residente

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServiceVisitaResidenteRepository interface {
	create(m *VisitaDomiciliaria) error
	getByID(id uint64) (*VisitaDomiciliaria, error)
	update(m *VisitaDomiciliaria) error
	delete(id uint64) error
	getAll() ([]*VisitaDomiciliaria, error)
	getAllWithFilters(filtros *FiltrosVisita) ([]*VisitaDomiciliaria, error)
	
	getByAlumnoID(alumnoID uint64) (*VisitaDomiciliaria, error)
	existsByAlumnoID(alumnoID uint64) (bool, error)
	getAlumnosPendientesVisitaPorConvocatoria(convocatoriaID uint64) ([]*AlumnoPendienteVisita, error)
	getTodosAlumnosPendientesVisita() ([]*AlumnoPendienteVisita, error)
	getEstadisticas(convocatoriaID *uint64) (*EstadisticasVisita, error)
	getEstadisticasPorEscuelaProfesional(convocatoriaID *uint64) ([]*EstadisticasPorEscuelaProfesional, error)
	getEstadisticasPorLugarProcedencia(convocatoriaID *uint64) ([]*EstadisticasPorLugarProcedencia, error)
	getAlumnosPendientesPorDepartamento(convocatoriaID uint64, departamento string) ([]*AlumnoPendienteVisitaPorDepartamento, error)
	getTodosAlumnosPorConvocatoria(convocatoriaID uint64) ([]*AlumnoCompleto, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServiceVisitaResidenteRepository {
	return newVisitaResidenteSqlServerRepository(db, txID)
}