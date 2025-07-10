package response_messages

import "github.com/jmoiron/sqlx"

type Message struct {
	db *sqlx.DB
}

func NewMsg(db *sqlx.DB) Message {
	return Message{
		db: db,
	}
}

func (m *Message) GetByCode(code int) (int, string, string) {
	srvMsg := NewMessageService(FactoryStorage(m.db))
	msg, _, err := srvMsg.GetMessageByID(code)
	if err != nil || msg == nil {
		return 11, "No se pudo cargar el mensaje", "ERROR"
	}

	return msg.Code, msg.MessageType, msg.Message

}
