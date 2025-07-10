package models

type Announcement struct {
	ID          int    `json:"id"`
	UserID      *int   `json:"user_id"`
	FechaInicio string `json:"fecha_inicio"`
	FechaFin    string `json:"fecha_fin"`
	Nombre      string `json:"nombre"`
	Activo      bool   `json:"activo"`
}
