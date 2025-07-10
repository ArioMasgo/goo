package models

import "github.com/asaskevich/govalidator"

type RequestDentistryConsultation struct {
	ConsultaAreaMedica ConsultationMedicalArea `json:"consulta_odontologia"`
	ExamenBucal        *BuccalTest             `json:"examen_bucal"`
	ConsultaBucal      *BuccalConsultation     `json:"consulta"`
	ProcedimientoBucal *BuccalProcedure        `json:"procedimiento"`
}

type BuccalTest struct {
	ID             string `json:"id" valid:"uuid,required"`
	OdontogramaIMG string `json:"odontograma_img"`
	CPOD           string `json:"cpod" valid:"-"`
	Observacion    string `json:"observacion" valid:"-"`
	IHOS           string `json:"ihos" valid:"-"`
	Comentarios    string `json:"comentarios" valid:"-"`
}

type BuccalConsultation struct {
	ID             string  `json:"id" valid:"uuid,required"`
	Relato         string  `json:"relato" valid:"-"`
	Diagnostico    string  `json:"diagnostico" valid:"-"`
	ExamenAuxiliar string  `json:"examen_auxiliar" valid:"-"`
	ExamenClinico  *string `json:"examen_clinico" valid:"-"`
	Tratamiento    string  `json:"tratamiento" valid:"-"`
	Indicaciones   string  `json:"indicaciones" valid:"-"`
	Comentarios    string  `json:"comentarios" valid:"-"`
}

type BuccalProcedure struct {
	ID                string `json:"id" valid:"uuid,required"`
	TipoProcedimiento string `json:"tipo_procedimiento" valid:"-"`
	Recibo            string `json:"recibo" valid:"-"`
	Costo             string `json:"costo" valid:"-"`
	FechaPago         string `json:"fecha_pago" valid:"-"`
	PiezaDental       string `json:"pieza_dental" valid:"-"`
	Comentarios       string `json:"comentarios" valid:"-"`
}

func (m *RequestDentistryConsultation) ValidDentistryConsultation() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
