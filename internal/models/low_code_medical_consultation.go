package models

import "github.com/asaskevich/govalidator"

type RequestMedicalConsultation struct {
	ConsultaMedicina      MedicalConsultation     `json:"consulta_medicina"`
	AtencionIntegralOtros *IntegralAttentionOther `json:"atencion_integral_otros"`
}

type GetAllMedicalConsultation struct {
	ID            string    `json:"id"`
	FechaConsulta string    `json:"fecha_consulta"`
	AreaAsignada  string    `json:"area_asignada"`
	AreaOrigen    string    `json:"area_origen"`
	Paciente      *Patients `json:"paciente"`
}

type MedicalConsultation struct {
	ID            string `json:"id"`
	IDPaciente    string `json:"paciente_id"`
	FechaConsulta string `json:"fecha_consulta"`
}

type IntegralAttentionOther struct {
	ID                     string `json:"id"`
	Fecha                  string `json:"fecha"`
	Hora                   string `json:"hora"`
	Edad                   string `json:"edad"`
	MotivoConsulta         string `json:"motivo_consulta"`
	TiempoEnfermedad       string `json:"tiempo_enfermedad"`
	Apetito                string `json:"apetito"`
	Sed                    string `json:"sed"`
	Suenio                 string `json:"suenio"`
	EstadoAnimo            string `json:"estado_animo"`
	Orina                  string `json:"orina"`
	Deposiciones           string `json:"deposiciones"`
	Temperatura            string `json:"temperatura"`
	PresionArterial        string `json:"presion_arterial"`
	FrecuenciaCardiaca     string `json:"frecuencia_cardiaca"`
	FrecuenciaRespiratoria string `json:"frecuencia_respiratoria"`
	Peso                   string `json:"peso"`
	Talla                  string `json:"talla"`
	IndiceMasaCorporal     string `json:"indice_masa_corporal"`
	Diagnostico            string `json:"diagnostico"`
	Tratamiento            string `json:"tratamiento"`
	ExamenesAxuliares      string `json:"examenes_axuliares"`
	Referencia             string `json:"referencia"`
	Observacion            string `json:"observacion"`
	NumeroRecibo           string `json:"numero_recibo"`
	Costo                  string `json:"costo"`
	FechaPago              string `json:"fecha_pago"`
}

type GeneralMedicineConsultation struct {
	ID            string `json:"id"`
	FechaHora     string `json:"fecha_hora"`
	Anamnesis     string `json:"anamnesis"`
	ExamenClinico string `json:"examen_clinico"`
	Indicaciones  string `json:"indicaciones"`
}

func (m *RequestMedicalConsultation) ValidMedicalConsultation() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
