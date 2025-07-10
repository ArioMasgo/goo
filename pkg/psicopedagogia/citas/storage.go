package citas

import (
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesCitaRepository interface {
	GetByID(id int) (*Cita, error)
	Create(c *Cita) (int64, error)
	Update(c *Cita) error
	Delete(id int) error
	GetAll() ([]*Cita, error)
	ExisteCitaEnFecha(fecha time.Time, dni string) (bool, error)
}

func FactoryStorage(db *sqlx.DB) ServicesCitaRepository {
	return NewCitaRepository(db)
}
