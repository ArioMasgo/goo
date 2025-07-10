package low_code_medical_area

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"dbu-api/pkg/medical_area"
	"dbu-api/pkg/submission"
	"fmt"
	"github.com/google/uuid"
)

func fetchConsultationAssignment(srvService *medical_area.ServerMedicalArea, id, txID string) (*models.ConsultationAssignment, error) {
	ConsultationAssignment, _, err := srvService.SrvConsultationAssignment.GetConsultationAssignmentByIDConsultation(id)
	if err != nil {
		logger.Error.Println(txID, " - failed to get ConsultationAssignment:", err)
		return nil, fmt.Errorf("failed to get ConsultationAssignment: %w", err)
	}
	if ConsultationAssignment == nil {
		return nil, fmt.Errorf("ConsultationAssignment not found")
	}
	return &models.ConsultationAssignment{
		IDConsulta:   ConsultationAssignment.IDConsulta,
		AreaAsignada: ConsultationAssignment.AreaAsignada,
	}, nil
}

func (s *NursingConsultationService) addAnnouncement(srvService *medical_area.ServerMedicalArea, paciente_id string, consulta_id string, area string) int {

	srv := submission.NewServerSubmission(s.db, s.usr, s.txID)
	announcement, _, err := srv.SrvConvocatorias.GetActiveSubmissions()

	if err != nil {
		return 0
	}

	if announcement == nil {
		return 0
	}

	idAnnouncementSignatures, _, _ := srvService.SrvAnnouncementSignatures.GetAnnouncementSignaturesByIDPatient(paciente_id)

	if idAnnouncementSignatures != nil {
		if area == "enfermería" {
			if idAnnouncementSignatures.FirmaEnfermeria != "" {
				return 0
			}
		}
		if area == "medicina" {
			if idAnnouncementSignatures.FirmaMedicina != "" {
				return 0
			}
		}

		_, _, err := srvService.SrvAnnouncementSignatures.UpdateAnnouncementSignatures(idAnnouncementSignatures.ID, announcement.ID, paciente_id, consulta_id, area)
		if err != nil {
			return 0
		}
		return 1
	}
	_, _, err = srvService.SrvAnnouncementSignatures.CreateAnnouncementSignatures(uuid.New().String(), announcement.ID, paciente_id, consulta_id, area)
	if err != nil {
		return 0
	}

	return 1

}

func (s *DentistryConsultationService) addAnnouncement(srvService *medical_area.ServerMedicalArea, paciente_id string, consulta_id string, area string) int {

	srv := submission.NewServerSubmission(s.db, s.usr, s.txID)
	announcement, _, err := srv.SrvConvocatorias.GetActiveSubmissions()

	if err != nil {
		return 0
	}

	if announcement == nil {
		return 0
	}

	idAnnouncementSignatures, _, _ := srvService.SrvAnnouncementSignatures.GetAnnouncementSignaturesByIDPatient(paciente_id)

	if idAnnouncementSignatures != nil {

		if idAnnouncementSignatures.FirmaOdontologia != "" {
			return 0
		}

		_, _, err := srvService.SrvAnnouncementSignatures.UpdateAnnouncementSignatures(idAnnouncementSignatures.ID, announcement.ID, paciente_id, consulta_id, area)
		if err != nil {
			return 0
		}
		return 1
	}
	_, _, err = srvService.SrvAnnouncementSignatures.CreateAnnouncementSignatures(uuid.New().String(), announcement.ID, paciente_id, consulta_id, area)
	if err != nil {
		return 0
	}

	return 1

}

func (s *MedicalConsultationService) addAnnouncement(srvService *medical_area.ServerMedicalArea, paciente_id string, consulta_id string, area string) int {

	srv := submission.NewServerSubmission(s.db, s.usr, s.txID)
	announcement, _, err := srv.SrvConvocatorias.GetActiveSubmissions()

	if err != nil {
		return 0
	}

	if announcement == nil {
		return 0
	}

	idAnnouncementSignatures, _, _ := srvService.SrvAnnouncementSignatures.GetAnnouncementSignaturesByIDPatient(paciente_id)

	if idAnnouncementSignatures != nil {

		if idAnnouncementSignatures.FirmaMedicina != "" {
			return 0
		}

		_, _, err := srvService.SrvAnnouncementSignatures.UpdateAnnouncementSignatures(idAnnouncementSignatures.ID, announcement.ID, paciente_id, consulta_id, area)
		if err != nil {
			return 0
		}
		return 1
	}
	_, _, err = srvService.SrvAnnouncementSignatures.CreateAnnouncementSignatures(uuid.New().String(), announcement.ID, paciente_id, consulta_id, area)
	if err != nil {
		return 0
	}

	return 1

}
