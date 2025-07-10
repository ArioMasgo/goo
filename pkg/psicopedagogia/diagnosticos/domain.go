package diagnosticos

import "time"

type Diagnostico struct {
	ID        int       `json:"id" db:"id"`
	Codigo    string    `json:"codigo" db:"codigo"`
	Nombre    string    `json:"nombre" db:"nombre"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
