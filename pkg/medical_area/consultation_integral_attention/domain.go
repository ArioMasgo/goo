package consultation_integral_attention

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type ConsultationIntegralAttention struct {
	ID                     string     `json:"id" db:"id"`
	IDConsulta             string     `json:"consulta_id" db:"consulta_id"`
	Fecha                  string     `json:"fecha" db:"fecha"`
	Hora                   string     `json:"hora" db:"hora"`
	Edad                   string     `json:"edad" db:"edad"`
	MotivoConsulta         string     `json:"motivo_consulta" db:"motivo_consulta"`
	TiempoEnfermedad       string     `json:"tiempo_enfermedad" db:"tiempo_enfermedad"`
	Apetito                string     `json:"apetito" db:"apetito"`
	Sed                    string     `json:"sed" db:"sed"`
	Suenio                 string     `json:"suenio" db:"suenio"`
	EstadoAnimo            string     `json:"estado_animo" db:"estado_animo"`
	Orina                  string     `json:"orina" db:"orina"`
	Deposiciones           string     `json:"deposiciones" db:"deposiciones"`
	Temperatura            string     `json:"temperatura" db:"temperatura"`
	PresionArterial        string     `json:"presion_arterial" db:"presion_arterial"`
	FrecuenciaCardiaca     string     `json:"frecuencia_cardiaca" db:"frecuencia_cardiaca"`
	FrecuenciaRespiratoria string     `json:"frecuencia_respiratoria" db:"frecuencia_respiratoria"`
	Peso                   string     `json:"peso" db:"peso"`
	Talla                  string     `json:"talla" db:"talla"`
	IndiceMasaCorporal     string     `json:"indice_masa_corporal" db:"indice_masa_corporal"`
	Diagnostico            string     `json:"diagnostico" db:"diagnostico"`
	Tratamiento            string     `json:"tratamiento" db:"tratamiento"`
	ExamenesAxuliares      string     `json:"examenes_axuliares" db:"examenes_axuliares"`
	Referencia             string     `json:"referencia" db:"referencia"`
	Observacion            string     `json:"observacion" db:"observacion"`
	NumeroRecibo           string     `json:"numero_recibo" db:"numero_recibo"`
	Costo                  string     `json:"costo" db:"costo"`
	FechaPago              string     `json:"fecha_pago" db:"fecha_pago"`
	IsDeleted              bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted            *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt              *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator            string     `json:"user_creator" db:"user_creator"`
	CreatedAt              *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt              *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}
type T struct {
	ConsultaGeneral struct {
	} `json:"consulta_general"`
}

func NewConsultationIntegralAttention(id, consulta_id, fecha, hora, edad, motivo_consulta, tiempo_enfermedad, apetito, sed, suenio, estado_animo, orina, deposiciones, temperatura, presion_arterial, frecuencia_cardiaca, frecuencia_respiratoria, peso, talla, indice_masa_corporal, diagnostico, tratamiento, examenes_axuliares, referencia, observacion, numero_recibo, costo, fecha_pago string) *ConsultationIntegralAttention {
	now := time.Now()
	return &ConsultationIntegralAttention{
		ID:                     id,
		IDConsulta:             consulta_id,
		Fecha:                  fecha,
		Hora:                   hora,
		Edad:                   edad,
		MotivoConsulta:         motivo_consulta,
		TiempoEnfermedad:       tiempo_enfermedad,
		Apetito:                apetito,
		Sed:                    sed,
		Suenio:                 suenio,
		EstadoAnimo:            estado_animo,
		Orina:                  orina,
		Deposiciones:           deposiciones,
		Temperatura:            temperatura,
		PresionArterial:        presion_arterial,
		FrecuenciaCardiaca:     frecuencia_cardiaca,
		FrecuenciaRespiratoria: frecuencia_respiratoria,
		Peso:                   peso,
		Talla:                  talla,
		IndiceMasaCorporal:     indice_masa_corporal,
		Diagnostico:            diagnostico,
		Tratamiento:            tratamiento,
		ExamenesAxuliares:      examenes_axuliares,
		Referencia:             referencia,
		Observacion:            observacion,
		NumeroRecibo:           numero_recibo,
		Costo:                  costo,
		FechaPago:              fecha_pago,
		IsDeleted:              false,
		CreatedAt:              &now,
		UpdatedAt:              &now,
	}
}

func (m *ConsultationIntegralAttention) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
