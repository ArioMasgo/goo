package payments_concept

import "time"

type ServicioMedicoConfig struct {
	ID                        string    `json:"id" db:"id" valid:"required"`
	Area                      string    `json:"area" db:"area" valid:"required,in(medicina|enfermeria|odontologia)"`
	TipoServicio              string    `json:"tipo_servicio" db:"tipo_servicio" valid:"required"`
	NombreServicio            string    `json:"nombre_servicio" db:"nombre_servicio" valid:"required"`
	RequierePago              bool      `json:"requiere_pago" db:"requiere_pago" valid:"required"`
	CodigoConcepto            *int      `json:"codigo_concepto,omitempty" db:"codigo_concepto" valid:"-"`
	ObligatorioEstudiante     bool      `json:"obligatorio_estudiante" db:"obligatorio_estudiante" valid:"-"`
	ObligatorioDocente        bool      `json:"obligatorio_docente" db:"obligatorio_docente" valid:"-"`
	ObligatorioAdministrativo bool      `json:"obligatorio_administrativo" db:"obligatorio_administrativo" valid:"-"`
	Estado                    bool      `json:"estado" db:"estado" valid:"-"`
	CreatedAt                 time.Time `json:"created_at" db:"created_at" valid:"-"`
	UpdatedAt                 time.Time `json:"updated_at" db:"updated_at" valid:"-"`
}

type DetallePagoTesoreria struct {
	NombreCompleto  string          `json:"nombre_completo"`
	Nombres         string          `json:"nombres"`
	ApellidoPaterno string          `json:"apellido_paterno"`
	ApellidoMaterno string          `json:"apellido_materno"`
	DNI             string          `json:"dni"`
	PagosRealizados []PagoTesoreria `json:"pagos_realizados"`
}

// PagosRealizados representa cada pago individual realizado
type PagoTesoreria struct {
	CodigoConcepto int     `json:"codconcepto"`
	ConceptoPagado string  `json:"concepto_pagado"`
	PrecioUnit     string  `json:"precio_unit"`
	Cantidad       string  `json:"cantidad"`
	ImportePagado  string  `json:"importe_pagado"`
	FechaPago      string  `json:"fecha_pago"`
	CodRecibo      *string `json:"cod_recibo"`
	CodReciboCanje *string `json:"cod_recibo_canje"`
	TipoPago       string  `json:"tipo_pago"`
	EsCanje        bool    `json:"es_canje"`
}

type PagosServicios struct {
	Recibo   string `json:"recibo" db:"recibo" valid:"required"`
	Servicio string `json:"servicio" db:"servicio" valid:"required"`
}
