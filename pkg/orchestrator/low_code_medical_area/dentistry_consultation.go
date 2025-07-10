package low_code_medical_area

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"dbu-api/pkg/medical_area"
	"dbu-api/pkg/medical_area/consultation_medical_area"
	"encoding/base64"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"path/filepath"
	"strings"
)

// Insertar una consulta de odontología en la base de datos

type DentistryConsultationService struct {
	db   *sqlx.DB
	usr  *models.User
	txID string
}

type PortsServerDentistryConsultation interface {
	CreateDentistryConsultationLowCode(m *models.RequestDentistryConsultation) (int, error)
	UpdateDentistryConsultationLowCode(m *models.RequestDentistryConsultation) (int, error)
	DeleteDentistryConsultationLowCode(id string) (int, error)
	GetDentistryConsultationByIdConsultationLowCode(id string) (*models.RequestDentistryConsultation, int, error)
	GetDentistryConsultationByDNILowCode(dni string) (*models.GetAllConsultationMedicalArea, int, error)
	GetDentistryConsultationByIdPatientLowCode(id string) ([]*models.RequestDentistryConsultation, int, error)
	GetAllDentistryConsultationLowCode() ([]*models.GetAllConsultationMedicalArea, error)
}

func NewDentistryConsultation(db *sqlx.DB, usr *models.User, txID string) PortsServerDentistryConsultation {
	return &DentistryConsultationService{db: db, usr: usr, txID: txID}
}

func (s *DentistryConsultationService) CreateDentistryConsultationLowCode(m *models.RequestDentistryConsultation) (int, error) {

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	dataDentistryConsultation, code, err := srv.SrvConsultationMedicalArea.CreateConsultationMedicalArea(m.ConsultaAreaMedica.ID, m.ConsultaAreaMedica.IDPaciente, m.ConsultaAreaMedica.FechaConsulta, "odontología")
	if err != nil {
		logger.Error.Printf("couldn't create dentistry consultation, error: %v", dataDentistryConsultation, err)
		return code, err
	}

	dataConsultationAssignment, code, err := srv.SrvConsultationAssignment.CreateConsultationAssignment("", m.ConsultaAreaMedica.ID, "odontología")
	if err != nil {
		logger.Error.Printf("couldn't create consultation assignment, error: %v", dataConsultationAssignment, err)
		return code, err
	}

	if m.ExamenBucal != nil {

		routeOdontogram, code, err := saveBuccalTestImage(m.ExamenBucal.OdontogramaIMG, m.ConsultaAreaMedica.IDPaciente, m.ExamenBucal.ID)
		if err != nil {
			logger.Error.Printf("couldn't save buccal test image, error: %v", code, err)
			return code, err
		}

		dataBuccalTest, code, err := srv.SrvDentistryConsultationBuccalTest.CreateBuccalTest(m.ExamenBucal.ID, m.ConsultaAreaMedica.ID, routeOdontogram, m.ExamenBucal.IHOS, m.ExamenBucal.CPOD, m.ExamenBucal.Observacion, m.ExamenBucal.Comentarios)
		if err != nil {
			logger.Error.Printf("couldn't create buccal test, error: %v", dataBuccalTest, err)
			return code, err
		}

		s.addAnnouncement(srv, m.ConsultaAreaMedica.IDPaciente, m.ConsultaAreaMedica.ID, "odontología")
	}

	if m.ConsultaBucal != nil {
		dataBuccalConsultation, code, err := srv.SrvDentistryConsultationBuccalConsultation.CreateBuccalConsultation(m.ConsultaBucal.ID, m.ConsultaAreaMedica.ID, m.ConsultaBucal.Relato, m.ConsultaBucal.Diagnostico, m.ConsultaBucal.ExamenAuxiliar, m.ConsultaBucal.ExamenClinico, m.ConsultaBucal.Tratamiento, m.ConsultaBucal.Indicaciones, m.ConsultaBucal.Comentarios)
		if err != nil {
			logger.Error.Printf("couldn't create buccal test, error: %v", dataBuccalConsultation, err)
			return code, err
		}
	}

	if m.ProcedimientoBucal != nil {
		dataBuccalProcedure, code, err := srv.SrvDentistryConsultationBuccalProcedure.CreateBuccalProcedure(m.ProcedimientoBucal.ID, m.ConsultaAreaMedica.ID, m.ProcedimientoBucal.TipoProcedimiento, m.ProcedimientoBucal.Recibo, m.ProcedimientoBucal.Costo, m.ProcedimientoBucal.FechaPago, m.ProcedimientoBucal.PiezaDental, m.ProcedimientoBucal.Comentarios)
		if err != nil {
			logger.Error.Printf("couldn't create buccal test, error: %v", dataBuccalProcedure, err)
			return code, err
		}
	}

	//getIDAnnouncementActive(srv, m.ConsultaAreaMedica.IDPaciente, m.ConsultaAreaMedica.ID, "odontología")

	return 20, nil
}

func (s *DentistryConsultationService) UpdateDentistryConsultationLowCode(m *models.RequestDentistryConsultation) (int, error) {

	code, err := s.DeleteDentistryConsultationLowCode(m.ConsultaAreaMedica.ID)

	if err != nil {
		logger.Error.Printf("couldn't delete dentistry consultation, error: %v", err)
		return code, err
	}

	code, err = s.CreateDentistryConsultationLowCode(m)

	if err != nil {
		logger.Error.Printf("couldn't create dentistry consultation, error: %v", err)
		return code, err
	}

	return 20, nil
}

func (s *DentistryConsultationService) DeleteDentistryConsultationLowCode(id string) (int, error) {

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	code, err := srv.SrvDentistryConsultationBuccalTest.DeleteBuccalTestByIDConsultation(id)
	if err != nil {
		logger.Error.Printf("couldn't delete buccal test, error: %v", id, err)
		return code, err
	}

	code, err = srv.SrvDentistryConsultationBuccalConsultation.DeleteBuccalConsultationByIDConsultation(id)
	if err != nil {
		logger.Error.Printf("couldn't delete buccal consultation, error: %v", id, err)
		return code, err
	}

	code, err = srv.SrvDentistryConsultationBuccalProcedure.DeleteBuccalProcedureByIDConsultation(id)
	if err != nil {
		logger.Error.Printf("couldn't delete buccal procedure, error: %v", id, err)
		return code, err
	}

	codeConsultationAssignment, err := srv.SrvConsultationAssignment.DeleteConsultationAssignmentByIDConsultation(id)
	if err != nil {
		logger.Error.Printf("couldn't delete consultation assignment, error: %v", err)
		return codeConsultationAssignment, err
	}

	code, err = srv.SrvConsultationMedicalArea.DeleteConsultationMedicalArea(id)
	if err != nil {
		logger.Error.Printf("couldn't delete dentistry consultation, error: %v", id, err)
		return code, err
	}

	return 20, nil
}

func (s *DentistryConsultationService) GetDentistryConsultationByIdConsultationLowCode(id string) (*models.RequestDentistryConsultation, int, error) {

	srvService := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	dentistryConsultation, code, err := srvService.SrvConsultationMedicalArea.GetConsultationMedicalAreaByID(id)
	if err != nil {
		logger.Error.Printf("couldn't get dentistry consultation by id consultation, error: %v", id, err)
		return nil, code, err
	}

	if dentistryConsultation == nil {
		return nil, StatusNotFound, nil
	}

	buccalTest, err := fetchBuccalTest(srvService, dentistryConsultation.ID, s.txID)
	if err != nil {
		return nil, code, err
	}

	buccalConsultation, err := fetchBuccalConsultation(srvService, dentistryConsultation.ID, s.txID)
	if err != nil {
		return nil, code, err
	}
	buccalProcedure, err := fetchBuccalProcedure(srvService, dentistryConsultation.ID, s.txID)
	if err != nil {
		return nil, code, err
	}

	data := &models.RequestDentistryConsultation{
		ConsultaAreaMedica: mapDentistryConsultation(dentistryConsultation),
		ExamenBucal:        buccalTest,
		ConsultaBucal:      buccalConsultation,
		ProcedimientoBucal: buccalProcedure,
	}

	return data, StatusSuccess, nil
}

func (s *DentistryConsultationService) GetDentistryConsultationByIdPatientLowCode(id string) ([]*models.RequestDentistryConsultation, int, error) {
	srvService := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	dentistryConsultations, code, err := srvService.SrvConsultationMedicalArea.GetConsultationMedicalAreasByPatientID(id)
	if err != nil {
		logger.Error.Printf("couldn't get dentistry consultation by id patient, error: %v", id, err)
		return nil, code, err
	}

	if dentistryConsultations == nil {
		return nil, StatusNotFound, nil
	}

	var result []*models.RequestDentistryConsultation

	for _, dentistryConsultation := range dentistryConsultations {

		buccalTest, err := fetchBuccalTest(srvService, dentistryConsultation.ID, s.txID)
		if err != nil {
			return nil, code, err
		}

		buccalConsultation, err := fetchBuccalConsultation(srvService, dentistryConsultation.ID, s.txID)
		if err != nil {
			return nil, code, err
		}
		buccalProcedure, err := fetchBuccalProcedure(srvService, dentistryConsultation.ID, s.txID)
		if err != nil {
			return nil, code, err
		}

		data := &models.RequestDentistryConsultation{
			ConsultaAreaMedica: mapDentistryConsultation(dentistryConsultation),
			ExamenBucal:        buccalTest,
			ConsultaBucal:      buccalConsultation,
			ProcedimientoBucal: buccalProcedure,
		}

		result = append(result, data)
	}

	return result, StatusSuccess, nil
}

func (s *DentistryConsultationService) GetDentistryConsultationByDNILowCode(dni string) (*models.GetAllConsultationMedicalArea, int, error) {
	srvService := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	consultations, code, err := srvService.SrvConsultationMedicalArea.GetConsultationMedicalAreaByPatientDNI(dni)

	if err != nil {
		logger.Error.Println(s.txID, " - failed to get all Dentistry Consultations:", err)
		return nil, code, fmt.Errorf("failed to get all Dentistry Consultations: %w", err)
	}
	if consultations == nil {
		return nil, code, nil
	}

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

		services, _, err := s.fetchServicesDentistry(srvService, consultation.ID)
		if err != nil {
			return nil, code, err
		}

		consultationResponse := models.ResponseConsultationMedicalArea{
			ID:            consultation.ID,
			FechaConsulta: consultation.FechaConsulta,
			AreaAsignada:  ConsultationAssignment.AreaAsignada,
			AreaOrigen:    consultation.AreaMedica,
			Servicios:     services,
		}

		result = append(result, consultationResponse)

	}

	data := &models.GetAllConsultationMedicalArea{
		Paciente:  patient,
		Consultas: result,
	}

	return data, code, nil
}

func (s *DentistryConsultationService) GetAllDentistryConsultationLowCode() ([]*models.GetAllConsultationMedicalArea, error) {
	srvService := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	consultations, err := srvService.SrvConsultationMedicalArea.GetAllConsultationMedicalArea()

	if err != nil {
		logger.Error.Println(s.txID, " - failed to get all dentistry Consultations:", err)
		return nil, fmt.Errorf("failed to get all dentistry Consultations: %w", err)
	}
	if consultations == nil {
		return nil, nil
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

	var result []*models.GetAllConsultationMedicalArea

	for _, consultation := range consultations {

		patient, err := fetchResposePatientInfo(srvService, consultation.IDPaciente, s.txID)
		if err != nil {
			return nil, err
		}

		ConsultationAssignment, err := fetchConsultationAssignment(srvService, consultation.ID, s.txID)
		if err != nil {
			return nil, err
		}

		services, _, err := s.fetchServicesDentistry(srvService, consultation.ID)
		if err != nil {
			return nil, err
		}

		consultationResponse := models.ResponseConsultationMedicalArea{
			ID:            consultation.ID,
			FechaConsulta: consultation.FechaConsulta,
			AreaAsignada:  ConsultationAssignment.AreaAsignada,
			AreaOrigen:    consultation.AreaMedica,
			Servicios:     services,
		}
		data := &models.GetAllConsultationMedicalArea{
			Paciente:  patient,
			Consultas: []models.ResponseConsultationMedicalArea{consultationResponse},
		}

		result = append(result, data)
	}

	return result, nil
}

func (s *DentistryConsultationService) fetchServicesDentistry(srvService *medical_area.ServerMedicalArea, id string) (string, int, error) {
	var services []string

	buccalTest, err := fetchBuccalTest(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if buccalTest != nil {
		services = append(services, "Examen Bucal")
	}

	buccalConsultation, err := fetchBuccalConsultation(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if buccalConsultation != nil {
		services = append(services, "Consulta Bucal")
	}

	buccalProcedure, err := fetchBuccalProcedure(srvService, id, s.txID)
	if err != nil {
		return "", StatusNotFound, err
	}
	if buccalProcedure != nil {
		services = append(services, "Procedimiento Bucal")
	}

	result := strings.Join(services, ", ")

	return result, 0, nil
}

func mapDentistryConsultation(dentistryConsultation *consultation_medical_area.ConsultationMedicalArea) models.ConsultationMedicalArea {
	return models.ConsultationMedicalArea{
		ID:            dentistryConsultation.ID,
		IDPaciente:    dentistryConsultation.IDPaciente,
		FechaConsulta: dentistryConsultation.FechaConsulta,
		AreaMedica:    &dentistryConsultation.AreaMedica,
	}
}

func fetchBuccalTest(srvService *medical_area.ServerMedicalArea, id, txID string) (*models.BuccalTest, error) {
	buccalTest, _, err := srvService.SrvDentistryConsultationBuccalTest.GetBuccalTestByIDConsultation(id)
	if err != nil {
		logger.Error.Println(txID, " - failed to get Buccal Test:", err)
		return nil, fmt.Errorf("failed to get Buccal Test: %w", err)
	}
	if buccalTest == nil {
		return nil, nil
	}

	buccalTest.OdontogramaIMG, err = getSVGAsBase64(buccalTest.OdontogramaIMG, buccalTest.ID)

	return &models.BuccalTest{
		ID:             buccalTest.ID,
		OdontogramaIMG: buccalTest.OdontogramaIMG,
		CPOD:           buccalTest.CPOD,
		Observacion:    buccalTest.Observacion,
		IHOS:           buccalTest.IHOS,
		Comentarios:    buccalTest.Comentarios,
	}, nil
}

func fetchBuccalConsultation(srvService *medical_area.ServerMedicalArea, id, txID string) (*models.BuccalConsultation, error) {
	buccalConsultation, _, err := srvService.SrvDentistryConsultationBuccalConsultation.GetBuccalConsultationByIDConsultation(id)
	if err != nil {
		logger.Error.Println(txID, " - failed to get Buccal Consultation:", err)
		return nil, fmt.Errorf("failed to get Buccal Consultation: %w", err)
	}
	if buccalConsultation == nil {
		return nil, nil
	}
	return &models.BuccalConsultation{
		ID:             buccalConsultation.ID,
		Relato:         buccalConsultation.Relato,
		Diagnostico:    buccalConsultation.Diagnostico,
		ExamenAuxiliar: buccalConsultation.ExamenAuxiliar,
		ExamenClinico:  buccalConsultation.ExamenClinico,
		Tratamiento:    buccalConsultation.Tratamiento,
		Indicaciones:   buccalConsultation.Indicaciones,
		Comentarios:    buccalConsultation.Comentarios,
	}, nil
}

func fetchBuccalProcedure(srvService *medical_area.ServerMedicalArea, id, txID string) (*models.BuccalProcedure, error) {
	buccalProcedure, _, err := srvService.SrvDentistryConsultationBuccalProcedure.GetBuccalProcedureByIDConsultation(id)
	if err != nil {
		logger.Error.Println(txID, " - failed to get Buccal Procedure:", err)
		return nil, fmt.Errorf("failed to get Buccal Procedure: %w", err)
	}
	if buccalProcedure == nil {
		return nil, nil
	}
	return &models.BuccalProcedure{
		ID:                buccalProcedure.ID,
		TipoProcedimiento: buccalProcedure.TipoProcedimiento,
		Recibo:            buccalProcedure.Recibo,
		Costo:             buccalProcedure.Costo,
		FechaPago:         buccalProcedure.FechaPago,
		PiezaDental:       buccalProcedure.PiezaDental,
		Comentarios:       buccalProcedure.Comentarios,
	}, nil
}

func saveBuccalTestImage(imageData string, pacienteID string, examenID string) (string, int, error) {
	basePath := "/storage/app/public/medical_area/odontogram"

	svgBytes, err := base64.StdEncoding.DecodeString(imageData)

	dirPath := filepath.Join(basePath, pacienteID)
	filePath := filepath.Join(dirPath, examenID+".svg")

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return "", StatusNotFound, fmt.Errorf("error creando directorio: %w", err)
		}
	}

	err = os.WriteFile(filePath, svgBytes, 0644)
	if err != nil {
		return "", StatusNotFound, fmt.Errorf("error guardando la imagen: %w", err)
	}

	fmt.Println("Archivo guardado en:", filePath)
	return dirPath, StatusSuccess, nil
}

func getSVGAsBase64(route string, name string) (string, error) {

	filePath := filepath.Join(route, name+".svg")

	svgBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error al leer el archivo SVG: %w", err)
	}

	base64SVG := base64.StdEncoding.EncodeToString(svgBytes)

	return base64SVG, nil
}
