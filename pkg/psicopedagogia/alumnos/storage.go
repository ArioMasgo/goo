package alumnos

import (
	"dbu-api/pkg/psicopedagogia/participantes"

	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesEstudianteRepository interface {
	getByDni(dni string) (*Estudiante, error)
	getNameByDni(dni string) (*BasicEstudiante, error)
	GetParticipanteByDNI(dni string, tipoParticipante string) (*participantes.Participante, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesEstudianteRepository {
	return newEstudianteSqlServerRepository(db, txID)
}
