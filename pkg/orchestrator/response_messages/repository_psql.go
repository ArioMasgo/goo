package response_messages

import (
	"database/sql"
	"dbu-api/internal/logger"

	"github.com/jmoiron/sqlx"
)

// mysql estructura de conexión a la BD de postgresql
type mysql struct {
	DB *sqlx.DB
}

func NewMessagePsqlRepository(db *sqlx.DB) *mysql {
	return &mysql{
		DB: db,
	}
}

// GetByID consulta un registro por su ID
func (s *mysql) GetByID(id int) (*ResponseMessage, error) {
	const sqlGetByID = `SELECT id, code, message, message_type, http_status, created_at, updated_at 
        FROM response_messages 
        WHERE code = ?`
	mdl := ResponseMessage{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(" - couldn't execute GetByID Message: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}
