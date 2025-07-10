package models

import "github.com/asaskevich/govalidator"

type RequestNursingConsultation struct {
	ConsultaAreaMedica ConsultationMedicalArea `json:"consulta_enfermeria"`
	RevisionRutina     RoutineReview           `json:"revision_rutina"`
	DatosAcompanante   *AccompanyingData       `json:"datos_acompanante"`
	Examenes           Exams                   `json:"examenes"`
}

type RoutineReview struct {
	ID                       string `json:"id" valid:"required"`
	FiebreUltimoQuinceDias   string `json:"fiebre_ultimo_quince_dias" valid:"-"`
	TosMasQuinceDias         string `json:"tos_mas_quince_dias" valid:"-"`
	SecrecionLesionGenitales string `json:"secrecion_lesion_genitales" valid:"-"`
	FechaUltimaRegla         string `json:"fecha_ultima_regla" valid:"-"`
	Comentarios              string `json:"comentarios" valid:"-"`
}

type AccompanyingData struct {
	ID               string `json:"id" valid:"required"`
	DNI              string `json:"dni" valid:"-"`
	NombresApellidos string `json:"nombres_apellidos" valid:"-"`
	Edad             string `json:"edad" valid:"-"`
}

type Exams struct {
	Vacunas                  []*Vaccine                   `json:"vacunas"`
	ExamenFisico             *PhysicalTest                `json:"examen_fisico"`
	ExamenLaboratorio        *LaboratoryTest              `json:"examen_laboratorio"`
	ExamenPreferencial       *PreferentialTest            `json:"examen_preferencial"`
	ExamenSexualidad         *SexualityTest               `json:"examen_sexualidad"`
	ExamenVisual             *VisualTest                  `json:"examen_visual"`
	TratamientoMedicamentoso []*MedicationTreatment       `json:"tratamiento_medicamentoso"`
	ProcedimientosRealizados []*PerformedProcedures       `json:"procedimientos_realizados"`
	AtencionIntegralOtros    *IntegralAttentionOther      `json:"consulta_general"`
	ConsultaBucal            *BuccalConsultation          `json:"consulta"`
	ConsultaMedicinaGeneral  *GeneralMedicineConsultation `json:"consulta_medicina_general"`
}

type Vaccine struct {
	ID          string `json:"id" valid:"required"`
	TipoVacuna  string `json:"tipo_vacuna" valid:"-"`
	FechaDosis  string `json:"fecha_dosis" valid:"-"`
	Comentarios string `json:"comentarios" valid:"-"`
}

type PhysicalTest struct {
	ID                    string `json:"id" valid:"required"`
	TallaPesos            string `json:"talla_peso" valid:"-"`
	PerimetroCintura      string `json:"perimetro_cintura" valid:"-"`
	IndiceMasaCorporalImg string `json:"indice_masa_corporal_img" valid:"-"`
	PresionArterial       string `json:"presion_arterial" valid:"-"`
	Comentarios           string `json:"comentarios" valid:"-"`
}

type LaboratoryTest struct {
	ID          string `json:"id" valid:"required"`
	Serologia   string `json:"serologia" valid:"-"`
	Bk          string `json:"bk" valid:"-"`
	Hemograma   string `json:"hemograma" valid:"-"`
	ExamenOrina string `json:"examen_orina" valid:"-"`
	Colesterol  string `json:"colesterol" valid:"-"`
	Glucosa     string `json:"glucosa" valid:"-"`
	Comentarios string `json:"comentarios" valid:"-"`
}

type PreferentialTest struct {
	ID                    string `json:"id" valid:"required"`
	AparatoRespiratorio   string `json:"aparato_respiratorio" valid:"-"`
	AparatoCardiovascular string `json:"aparato_cardiovascular" valid:"-"`
	AparatoDigestivo      string `json:"aparato_digestivo" valid:"-"`
	AparatoGenitourinario string `json:"aparato_genitourinario" valid:"-"`
	Papanicolau           string `json:"papanicolau" valid:"-"`
	ExamenMama            string `json:"examen_mama" valid:"-"`
	Comentarios           string `json:"comentarios" valid:"-"`
}

type SexualityTest struct {
	ID                    string `json:"id" valid:"required"`
	ActividadSexual       string `json:"actividad_sexual" valid:"-"`
	PlanificacionFamiliar string `json:"planificacion_familiar" valid:"-"`
	Comentarios           string `json:"comentarios" valid:"-"`
}

type VisualTest struct {
	ID            string `json:"id" valid:"required"`
	OjoDerecho    string `json:"ojo_derecho" valid:"-"`
	OjoIzquierdo  string `json:"ojo_izquierdo" valid:"-"`
	PresionOcular string `json:"presion_ocular" valid:"-"`
	Comentarios   string `json:"comentarios" valid:"-"`
}

type MedicationTreatment struct {
	ID                        string  `json:"id" valid:"required"`
	NombreGenericoMedicamento string  `json:"nombre_generico_medicamento"`
	ViaAdministracion         string  `json:"via_administracion"`
	HoraAplicacion            string  `json:"hora_aplicacion"`
	ResponsableAtencion       string  `json:"responsable_atencion"`
	Observaciones             string  `json:"observaciones"`
	AreaSolicitante           *string `json:"area_solicitante"`
	EspecialistaSolicitante   *string `json:"especialista_solicitante"`
}

type PerformedProcedures struct {
	ID                      string  `json:"id" valid:"required"`
	Procedimiento           string  `json:"procedimiento"`
	NumeroRecibo            string  `json:"numero_recibo"`
	Comentarios             string  `json:"comentarios"`
	Costo                   string  `json:"costo"`
	FechaPago               string  `json:"fecha_pago"`
	AreaSolicitante         *string `json:"area_solicitante"`
	EspecialistaSolicitante *string `json:"especialista_solicitante"`
}

type TypesVaccines struct {
	ID            string `json:"id"`
	Nombre        string `json:"nombre"`
	Estado        bool   `json:"estado"`
	DuracionMeses string `json:"duracion_meses"`
}

type TypesVaccineRequired struct {
	Nombre         string `json:"nombre"`
	FechaRequerida string `json:"fecha_requerida"`
	Requerido      bool   `json:"requerido"`
}

func (m *RequestNursingConsultation) ValidNursingConsultation() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
