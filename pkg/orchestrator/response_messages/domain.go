package response_messages

import (
	"time"
)

type ResponseMessage struct {
	ID          int64     `db:"id" json:"id"`
	Code        int       `db:"code" json:"code"`
	Message     string    `db:"message" json:"message"`
	MessageType string    `db:"message_type" json:"message_type" valid:"in(ERROR|WARNING|SUCCESS)"`
	HTTPStatus  int       `db:"http_status" json:"http_status"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
