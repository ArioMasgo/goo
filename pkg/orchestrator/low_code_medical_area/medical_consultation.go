package low_code_medical_area

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"dbu-api/pkg/medical_area"
	"dbu-api/pkg/medical_area/medical_consultation"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

// Insertar una consulta de medicina en la base de datos

type MedicalConsultationService struct {
	db   *sqlx.DB
	usr  *models.User
	txID string
}

type PortsServerMedicalConsultation interface {
	CreateMedicalConsultationLowCode(m *models.RequestMedicalConsultation) (int, error)
	UpdateMedicalConsultationLowCode(m *models.RequestMedicalConsultation) (int, error)
	DeleteMedicalConsultationLowCode(id string) (int, error)
	GetMedicalConsultationByIdConsultationLowCode(id string) (*models.RequestMedicalConsultation, int, error)
	GetMedicalConsultationByIdPatientLowCode(id string) ([]*models.RequestMedicalConsultation, int, error)
	GetMedicalConsultationByDNILowCode(dni string) (*models.GetAllConsultationMedicalArea, int, error)
	GetAllMedicalConsultationLowCode() ([]*models.GetAllMedicalConsultation, error)
}

func NewMedicalConsultation(db *sqlx.DB, usr *models.User, txID string) PortsServerMedicalConsultation {
	return &MedicalConsultationService{db: db, usr: usr, txID: txID}
}

func (s *MedicalConsultationService) CreateMedicalConsultationLowCode(m *models.RequestMedicalConsultation) (int, error) {

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	dataMedicalConsultation, code, err := srv.SrvMedicalConsultation.CreateMedicalConsultation(m.ConsultaMedicina.ID, m.ConsultaMedicina.IDPaciente, m.ConsultaMedicina.FechaConsulta, "medicina")
	if err != nil {
		logger.Error.Printf("couldn't create medical consultation, error: %v", dataMedicalConsultation, err)
		return code, err
	}

	s.addAnnouncement(srv, m.ConsultaMedicina.IDPaciente, m.ConsultaMedicina.ID, "medicina")

	//dataIntegralAttentionOther, code, err := srv.SrvConsultationIntegralAttention.CreateConsultationIntegralAttention(m.AtencionIntegralOtros.ID, m.ConsultaMedicina.ID, m.AtencionIntegralOtros.FechaHora, m.AtencionIntegralOtros.PresionArterial, m.AtencionIntegralOtros.SaturacionOxigeno, m.AtencionIntegralOtros.PulsoFrecuenciaCardiaca, m.AtencionIntegralOtros.NumeroRecibo, m.AtencionIntegralOtros.TemperaturaCorporal, m.AtencionIntegralOtros.Anamnesis, m.AtencionIntegralOtros.ExamenClinico, m.AtencionIntegralOtros.Indicaciones)
	//if err != nil {
	//	logger.Error.Printf("couldn't create integral attention other, error: %v", dataIntegralAttentionOther, err)
	//	return code, err
	//}

	return 20, nil
}

func (s *MedicalConsultationService) UpdateMedicalConsultationLowCode(m *models.RequestMedicalConsultation) (int, error) {

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	dataMedicalConsultation, code, err := srv.SrvMedicalConsultation.UpdateMedicalConsultation(m.ConsultaMedicina.ID, m.ConsultaMedicina.IDPaciente, m.ConsultaMedicina.FechaConsulta, "medicina")
	if err != nil {
		logger.Error.Printf("couldn't update medical consultation, error: %v", dataMedicalConsultation, err)
		return code, err
	}

	s.addAnnouncement(srv, m.ConsultaMedicina.IDPaciente, m.ConsultaMedicina.ID, "medicina")

	//dataIntegralAttentionOther, code, err := srv.SrvMedicalConsultationIntegralAttentionOther.UpdateIntegralAttentionOther(m.AtencionIntegralOtros.ID, m.ConsultaMedicina.IDPaciente, m.AtencionIntegralOtros.FechaHora, m.AtencionIntegralOtros.PresionArterial, m.AtencionIntegralOtros.SaturacionOxigeno, m.AtencionIntegralOtros.PulsoFrecuenciaCardiaca, m.AtencionIntegralOtros.NumeroRecibo, m.AtencionIntegralOtros.TemperaturaCorporal, m.AtencionIntegralOtros.Anamnesis, m.AtencionIntegralOtros.ExamenClinico, m.AtencionIntegralOtros.Indicaciones)
	//if err != nil {
	//	logger.Error.Printf("couldn't update integral attention other, error: %v", dataIntegralAttentionOther, err)
	//	return code, err
	//}

	return 20, nil

}

func (s *MedicalConsultationService) DeleteMedicalConsultationLowCode(id string) (int, error) {

	//srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	//code, err := srv.SrvMedicalConsultationIntegralAttentionOther.DeleteIntegralAttentionOtherByIDConsultation(id)
	//if err != nil {
	//	logger.Error.Printf("couldn't delete integral attention other test, error: %v", id, err)
	//	return code, err
	//}

	//code, err = srv.SrvMedicalConsultation.DeleteMedicalConsultation(id)
	//if err != nil {
	//	logger.Error.Printf("couldn't delete medical consultation, error: %v", id, err)
	//	return code, err
	//}

	return 20, nil
}

func (s *MedicalConsultationService) GetMedicalConsultationByIdConsultationLowCode(id string) (*models.RequestMedicalConsultation, int, error) {
	srvService := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	medicalConsultation, code, err := srvService.SrvMedicalConsultation.GetMedicalConsultationByID(id)
	if err != nil {
		logger.Error.Printf("couldn't get medical consultation by id consultation, error: %v", id, err)
		return nil, code, err
	}

	if medicalConsultation == nil {
		return nil, StatusNotFound, nil
	}

	integralAttentionOther, err := fetchIntegralAttentionOther(srvService, medicalConsultation.ID, s.txID)
	if err != nil {
		return nil, code, err
	}

	data := &models.RequestMedicalConsultation{
		ConsultaMedicina:      mapMedicalConsultation(medicalConsultation),
		AtencionIntegralOtros: integralAttentionOther,
	}

	return data, StatusSuccess, nil
}

func (s *MedicalConsultationService) GetMedicalConsultationByIdPatientLowCode(id string) ([]*models.RequestMedicalConsultation, int, error) {
	srvService := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	medicalConsultations, code, err := srvService.SrvMedicalConsultation.GetMedicalConsultationByIDPatient(id)
	if err != nil {
		logger.Error.Printf("couldn't get medical consultation by id patient, error: %v", id, err)
		return nil, code, err
	}

	if medicalConsultations == nil {
		return nil, StatusNotFound, nil
	}

	var result []*models.RequestMedicalConsultation

	for _, medicalConsultation := range medicalConsultations {

		integralAttentionOther, err := fetchIntegralAttentionOther(srvService, medicalConsultation.ID, s.txID)
		if err != nil {
			return nil, code, err
		}
		data := &models.RequestMedicalConsultation{
			ConsultaMedicina:      mapMedicalConsultation(medicalConsultation),
			AtencionIntegralOtros: integralAttentionOther,
		}

		result = append(result, data)
	}

	return result, StatusSuccess, nil
}

func (s *MedicalConsultationService) GetMedicalConsultationByDNILowCode(dni string) (*models.GetAllConsultationMedicalArea, int, error) {
	srvService := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	consultations, code, err := srvService.SrvConsultationMedicalArea.GetConsultationMedicalAreaByPatientDNI(dni)

	if err != nil {
		logger.Error.Println(s.txID, " - failed to get all Medical Consultations:", err)
		return nil, code, fmt.Errorf("failed to get all Medical Consultations: %w", err)
	}
	if consultations == nil {
		return nil, code, nil
	}

	//nursingConsultations, _, err := srvService.SrvConsultationAssignment.GetConsultationAssignmentByArea("enfermería")
	//
	//if err != nil {
	//	logger.Error.Println(s.txID, " - failed to get all Nursing Consultations:", err)
	//	return nil, fmt.Errorf("failed to get all Nursing Consultations: %w", err)
	//}
	//if nursingConsultations == nil {
	//	return nil, nil
	//}

	var result []models.ResponseConsultationMedicalArea

	patient, err := fetchResposePatientInfo(srvService, dni, s.txID)
	if err != nil {
		return nil, code, err
	}

	for _, consultation := range consultations {

		ConsultationAssignment, err := fetchConsultationAssignment(srvService, consultation.ID, s.txID)
		if err != nil {
			return nil, code, err
		}

		//services, _, err := s.fetchServicesMedical(srvService, consultation.ID)
		//if err != nil {
		//	return nil, code, err
		//}

		consultationResponse := models.ResponseConsultationMedicalArea{
			ID:            consultation.ID,
			FechaConsulta: consultation.FechaConsulta,
			AreaAsignada:  ConsultationAssignment.AreaAsignada,
			AreaOrigen:    consultation.AreaMedica,
			Servicios:     "",
		}

		result = append(result, consultationResponse)

	}

	data := &models.GetAllConsultationMedicalArea{
		Paciente:  patient,
		Consultas: result,
	}

	return data, code, nil
}

func (s *MedicalConsultationService) fetchServicesMedical(srvService *medical_area.ServerMedicalArea, id string) (string, int, error) {
	var services []string

	vaccines, err := fetchVaccines(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if vaccines != nil {
		services = append(services, "Vacunas")
	}

	physicalTest, err := fetchPhysicalTest(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if physicalTest != nil {
		services = append(services, "Examen Físico")
	}

	laboratoryTest, err := fetchLaboratoryTest(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if laboratoryTest != nil {
		services = append(services, "Examen Laboratorio")
	}

	preferentialTest, err := fetchPreferentialTest(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if preferentialTest != nil {
		services = append(services, "Examen Preferencial")
	}

	sexualityTest, err := fetchSexualityTest(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if sexualityTest != nil {
		services = append(services, "Examen Sexualidad")
	}

	visualTest, err := fetchVisualTest(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if visualTest != nil {
		services = append(services, "Examen Visual")
	}

	medicationTreatment, err := fetchMedicationTreatment(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if medicationTreatment != nil {
		services = append(services, "Tratamiento Medicamentoso")
	}

	performedProcedures, err := fetchPerformedProcedures(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if performedProcedures != nil {
		services = append(services, "Procedimientos Realizados")
	}

	consultationIntegralAttention, err := fetchConsultationIntegralAttention(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if consultationIntegralAttention != nil {
		services = append(services, "Atención Integral")
	}

	result := strings.Join(services, ", ")

	return result, StatusSuccess, nil
}

func (s *MedicalConsultationService) GetAllMedicalConsultationLowCode() ([]*models.GetAllMedicalConsultation, error) {
	srvService := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	medicalConsultations, err := srvService.SrvMedicalConsultation.GetAllMedicalConsultation()
	if err != nil {
		logger.Error.Println(s.txID, " - failed to get all Medical Consultations:", err)
		return nil, fmt.Errorf("failed to get all Medical Consultations: %w", err)
	}
	if medicalConsultations == nil {
		return nil, nil
	}

	var result []*models.GetAllMedicalConsultation

	for _, medicalConsultation := range medicalConsultations {

		patient, err := fetchPatient(srvService, medicalConsultation.IDPaciente, s.txID)
		if err != nil {
			return nil, err
		}

		ConsultationAssignment, err := fetchConsultationAssignment(srvService, medicalConsultation.ID, s.txID)
		if err != nil {
			return nil, err
		}

		data := &models.GetAllMedicalConsultation{
			ID:            medicalConsultation.ID,
			FechaConsulta: medicalConsultation.FechaConsulta,
			AreaAsignada:  ConsultationAssignment.AreaAsignada,
			AreaOrigen:    medicalConsultation.AreaMedica,
			Paciente:      patient,
		}

		result = append(result, data)
	}

	return result, nil
}

func mapMedicalConsultation(medicalConsultation *medical_consultation.MedicalConsultation) models.MedicalConsultation {
	return models.MedicalConsultation{
		ID:            medicalConsultation.ID,
		IDPaciente:    medicalConsultation.IDPaciente,
		FechaConsulta: medicalConsultation.FechaConsulta,
	}
}

func fetchIntegralAttentionOther(srvService *medical_area.ServerMedicalArea, id, txID string) (*models.IntegralAttentionOther, error) {
	//integralAttentionOther, _, err := srvService.SrvMedicalConsultationIntegralAttentionOther.GetIntegralAttentionOtherByIDConsultation(id)
	//if err != nil {
	//	logger.Error.Println(txID, " - failed to get Integral Attention Other:", err)
	//	return nil, fmt.Errorf("failed to get Integral Attention Other: %w", err)
	//}
	//if integralAttentionOther == nil {
	//	return nil, nil
	//}
	//return &models.IntegralAttentionOther{
	//	ID:                      integralAttentionOther.ID,
	//	FechaHora:               integralAttentionOther.FechaHora,
	//	PresionArterial:         integralAttentionOther.PresionArterial,
	//	SaturacionOxigeno:       integralAttentionOther.SaturacionOxigeno,
	//	PulsoFrecuenciaCardiaca: integralAttentionOther.PulsoFrecuenciaCardiaca,
	//	NumeroRecibo:            integralAttentionOther.NumeroRecibo,
	//	TemperaturaCorporal:     integralAttentionOther.TemperaturaCorporal,
	//	Anamnesis:               integralAttentionOther.Anamnesis,
	//	ExamenClinico:           integralAttentionOther.ExamenClinico,
	//	Indicaciones:            integralAttentionOther.Indicaciones,
	//}, nil
	return nil, nil
}
