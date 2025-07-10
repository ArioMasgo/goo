package models

type HeaderMedicalAreaExcel struct {
	Frame string `json:"frame"`
	Title string `json:"title"`
	Area  string `json:"area"`
}

type ConsultationPatientsMedicalAreaExcel struct {
	ID                  string `json:"id" db:"id" valid:"uuid,required"`
	FechaConsulta       string `json:"fecha_consulta" db:"fecha_consulta"`
	TipoPersona         string `json:"tipo_persona" db:"tipo_persona"`
	DNI                 string `json:"dni" db:"dni"`
	NombreCompleto      string `json:"nombre_completo" db:"nombre_completo"`
	Sexo                string `json:"sexo" db:"sexo"`
	FechaNacimiento     string `json:"fecha_nacimiento" db:"fecha_nacimiento"`
	Ocupacion           string `json:"ocupacion" db:"ocupacion"`
	NumeroCelular       string `json:"numero_celular" db:"numero_celular"`
	Servicios           string `json:"servicios" db:"servicios"`
	TipoProcedimiento   string `json:"tipo_procedimiento" db:"tipo_procedimiento"`
	Recibo              string `json:"recibo" db:"recibo"`
	Costo               string `json:"costo" db:"costo"`
	FechaPago           string `json:"fecha_pago" db:"fecha_pago"`
	PiezaDental         string `json:"pieza_dental" db:"pieza_dental"`
	CodigoSGA           string `json:"codigo_sga" db:"codigo_sga"`
	Procedencia         string `json:"procedencia" db:"procedencia"`
	DireccionResidencia string `json:"direccion_residencia" db:"direccion_residencia"`
	Procedimiento       string `json:"procedimiento" db:"procedimiento"`
	NumeroRecibo        string `json:"numero_recibo" db:"numero_recibo"`
	Monto               string `json:"monto" db:"monto"`
}

type PerformedProceduresExcel struct {
	ID                 string `json:"id" db:"id" valid:"uuid,required"`
	FechaConsulta      string `json:"fecha_consulta" db:"fecha_consulta"`
	TipoPersona        string `json:"tipo_persona" db:"tipo_persona"`
	EscuelaProfesional string `json:"escuela_profesional" db:"escuela_profesional"`
	Sexo               string `json:"sexo" db:"sexo"`
	TipoProcedimiento  string `json:"tipo_procedimiento" db:"tipo_procedimiento"`
}

type ConsultationIntegralAttentionExcel struct {
	ID                 string `json:"id" db:"id" valid:"uuid,required"`
	FechaConsulta      string `json:"fecha_consulta" db:"fecha_consulta"`
	TipoPersona        string `json:"tipo_persona" db:"tipo_persona"`
	EscuelaProfesional string `json:"escuela_profesional" db:"escuela_profesional"`
	Sexo               string `json:"sexo" db:"sexo"`
}
