package payments_concept

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	TxID string
}

func newPaymentConceptSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s *sqlserver) search(area, tipoServicio, nombreServicio string) (*ServicioMedicoConfig, error) {
	ml := &ServicioMedicoConfig{}
	const query = `SELECT id, area, tipo_servicio, nombre_servicio, requiere_pago, codigo_concepto, obligatorio_estudiante, obligatorio_docente, obligatorio_administrativo, estado, created_at, updated_at FROM servicios_medicos_config
	WHERE area = ? AND tipo_servicio = ? AND nombre_servicio = ?`
	err := s.DB.Get(ml, query, area, tipoServicio, nombreServicio)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return ml, nil
}

func (s *sqlserver) searchPaymentProcedureOdontologia(recibo, nombreServicio string) ([]*PagosServicios, error) {
	var ml []*PagosServicios
	const query = `SELECT recibo, tipo_procedimiento servicio FROM odontologia_procedimiento
	WHERE recibo = ? and tipo_procedimiento = ?`
	err := s.DB.Select(&ml, query, recibo, nombreServicio)
	if err != nil {
		return nil, err
	}

	return ml, nil
}
