package response_messages

import (
	"github.com/jmoiron/sqlx"
)

type ServicesMessageRepository interface {
	GetByID(id int) (*ResponseMessage, error)
}

func FactoryStorage(db *sqlx.DB) ServicesMessageRepository {
	return NewMessagePsqlRepository(db)
}
