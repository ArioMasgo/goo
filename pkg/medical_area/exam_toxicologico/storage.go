package exam_toxicologico

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)
type ServicesRegistroToxicologicoRepository interface {
	create(m *RegistroToxicologico) error
	update(m *RegistroToxicologico) error
	delete(id int64) error
	getByID(id int64) (*RegistroToxicologico, error)
	getByAlumnoAndConvocatoria(alumnoID int64, convocatoriaID int64) (*RegistroToxicologico, error)
	getAll() ([]*RegistroToxicologico, error)
	getEstadosByConvocatoria(convocatoriaID int64) ([]*EstadoToxicologicoConvocatoria, error)
	existsByAlumnoAndConvocatoria(alumnoID int64, convocatoriaID int64) (bool, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesRegistroToxicologicoRepository {
	return newRegistroToxicologicoSqlServerRepository(db, txID)
}
