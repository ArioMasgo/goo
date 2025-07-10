package models

type ConsultationMedicalArea struct {
	ID            string  `json:"id" valid:"required"`
	IDPaciente    string  `json:"paciente_id" valid:"required"`
	FechaConsulta string  `json:"fecha_consulta" valid:"required"`
	AreaMedica    *string `json:"area_medica" valid:"-"`
}

type GetAllConsultationMedicalArea struct {
	Paciente  *ResponsePatientInfo              `json:"paciente"`
	Consultas []ResponseConsultationMedicalArea `json:"consultas"`
}

type ResponseConsultationMedicalArea struct {
	ID            string `json:"id"`
	FechaConsulta string `json:"fecha_consulta"`
	AreaAsignada  string `json:"area_asignada"`
	AreaOrigen    string `json:"area_origen"`
	Servicios     string `json:"servicios"`
}
