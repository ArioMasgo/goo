package patients

type QueryParamsSearchPatients struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type RequestSearchPatients struct {
	Dni       string `json:"dni" valid:"-"`
	Nombres   string `json:"nombres" valid:"-"`
	Apellidos string `json:"apellidos" valid:"-"`
}
