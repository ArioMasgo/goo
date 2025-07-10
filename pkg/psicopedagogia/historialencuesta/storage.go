package historialencuesta

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesHistorialEncuestaRepository interface {
	Create(historial *HistorialEncuesta) (int, error)
	GetHistorialFiltered(dni, apellido, fechaInicio, fechaFin string) ([]Historial, error)
	HasStudentResponded(dni string, encuestaID int) (bool, error)
	GetLatestHistorial(limit, offset int) ([]Historial, error)
	KeyUrlExists(keyUrl string) (bool, error)
}

func FactoryStorage(db *sqlx.DB) ServicesHistorialEncuestaRepository {
	return NewHistorialRepository(db)
}
