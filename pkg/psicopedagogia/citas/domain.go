package citas

import "time"

type Cita struct {
	ID          int       `json:"id" db:"id"`
	DNI         string    `json:"dni" db:"dni"`
	Nombre      string    `json:"nombre" db:"nombre"`
	Apellido    string    `json:"apellido" db:"apellido"`
	Facultad    string    `json:"facultad" db:"facultad"`
	FechaInicio time.Time `json:"fecha_inicio" db:"fecha_inicio"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	FechaFin    time.Time `json:"fecha_fin" db:"fecha_fin"`
	Estado      string    `json:"estado" db:"estado"`
}
