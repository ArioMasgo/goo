package diagnosticos

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesDiagnosticoRepository interface {
	GetByID(id int) (*Diagnostico, error)
	Create(d *Diagnostico) (int64, error)
	Update(d *Diagnostico) error
	Delete(id int) error
	GetAll() ([]*Diagnostico, error)
}

func FactoryStorage(db *sqlx.DB) ServicesDiagnosticoRepository {
	return NewDiagnosticoRepository(db)
}
