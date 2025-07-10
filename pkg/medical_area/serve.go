package medical_area

import (
	"dbu-api/internal/models"
	"dbu-api/pkg/medical_area/announcement_signatures"
	"dbu-api/pkg/medical_area/consultation_assignment"
	"dbu-api/pkg/medical_area/consultation_integral_attention"
	"dbu-api/pkg/medical_area/consultation_medical_area"
	"dbu-api/pkg/medical_area/dentistry_consultation"
	"dbu-api/pkg/medical_area/dentistry_consultation_buccal_consultation"
	"dbu-api/pkg/medical_area/dentistry_consultation_buccal_procedure"
	"dbu-api/pkg/medical_area/dentistry_consultation_buccal_test"
	"dbu-api/pkg/medical_area/dentistry_consultation_odontogram_review"
	"dbu-api/pkg/medical_area/medical_consultation"
	"dbu-api/pkg/medical_area/medical_general_medicine_consultation"
	"dbu-api/pkg/medical_area/nursing_consultation_accompanying_data"
	"dbu-api/pkg/medical_area/nursing_consultation_laboratory_test"
	"dbu-api/pkg/medical_area/nursing_consultation_medication_treatment"
	"dbu-api/pkg/medical_area/nursing_consultation_performed_procedures"
	"dbu-api/pkg/medical_area/nursing_consultation_physical_test"
	"dbu-api/pkg/medical_area/nursing_consultation_preferential_test"
	"dbu-api/pkg/medical_area/nursing_consultation_routine_review"
	"dbu-api/pkg/medical_area/nursing_consultation_sexuality_test"
	"dbu-api/pkg/medical_area/nursing_consultation_vaccine"
	"dbu-api/pkg/medical_area/nursing_consultation_visual_test"
	"dbu-api/pkg/medical_area/patient_background"
	"dbu-api/pkg/medical_area/patients"
	"dbu-api/pkg/medical_area/exam_toxicologico"
	paymentsconcept "dbu-api/pkg/medical_area/payments"
	"github.com/jmoiron/sqlx"
)

type ServerMedicalArea struct {
	SrvExamToxicologico                        exam_toxicologico.PortsServerRegistroToxicologico
	SrvPatient                                 patients.PortsServerPatients
	SrvPatientBackground                       patient_background.PortsServerPatientBackground
	SrvConsultationMedicalArea                 consultation_medical_area.PortsServerConsultationMedicalArea
	SrvNursingConsultationRoutineReview        nursing_consultation_routine_review.PortsServerRoutineReview
	SrvNursingConsultationAccompanyingData     nursing_consultation_accompanying_data.PortsServerAccompanyingData
	SrvNursingConsultationVaccine              nursing_consultation_vaccine.PortsServerVaccine
	SrvNursingConsultationPhysicalTest         nursing_consultation_physical_test.PortsServerPhysicalTest
	SrvNursingConsultationLaboratoryTest       nursing_consultation_laboratory_test.PortsServerLaboratoryTest
	SrvNursingConsultationPreferentialTest     nursing_consultation_preferential_test.PortsServerPreferentialTest
	SrvNursingConsultationSexualityTest        nursing_consultation_sexuality_test.PortsServerSexualityTest
	SrvNursingConsultationVisualTest           nursing_consultation_visual_test.PortsServerVisualTest
	SrvNursingConsultationMedicationTreatment  nursing_consultation_medication_treatment.PortsServerMedicationTreatment
	SrvNursingConsultationPerformedProcedures  nursing_consultation_performed_procedures.PortsServerPerformedProcedures
	SrvDentistryConsultation                   dentistry_consultation.PortsServerDentistryConsultation
	SrvDentistryConsultationBuccalTest         dentistry_consultation_buccal_test.PortsServerBuccalTest
	SrvDentistryConsultationOdontogramReview   dentistry_consultation_odontogram_review.PortsServerOdontogramReview
	SrvDentistryConsultationBuccalConsultation dentistry_consultation_buccal_consultation.PortsServerBuccalConsultation
	SrvDentistryConsultationBuccalProcedure    dentistry_consultation_buccal_procedure.PortsServerBuccalProcedure
	SrvMedicalConsultation                     medical_consultation.PortsServerMedicalConsultation
	SrvConsultationIntegralAttention           consultation_integral_attention.PortsServerConsultationIntegralAttention
	SrvAnnouncementSignatures                  announcement_signatures.PortsServerAnnouncementSignatures
	SrvConsultationAssignment                  consultation_assignment.PortsServerConsultationAssignment
	SrvMedicalGeneralMedicineConsultation      medical_general_medicine_consultation.PortsServerGeneralMedicineConsultation
	SrvPaymentsConcept                         paymentsconcept.PortsServerPaymentConcept
}

func NewServerMedicalArea(db *sqlx.DB, usr *models.User, txID string) *ServerMedicalArea {
	return &ServerMedicalArea{
		SrvExamToxicologico:                        exam_toxicologico.NewRegistroToxicologicoService(exam_toxicologico.FactoryStorage(db, txID), usr, txID),
		SrvPatient:                                 patients.NewPatientsService(patients.FactoryStorage(db, txID), usr, txID),
		SrvPatientBackground:                       patient_background.NewPatientBackgroundService(patient_background.FactoryStorage(db, txID), usr, txID),
		SrvConsultationMedicalArea:                 consultation_medical_area.NewConsultationMedicalAreaService(consultation_medical_area.FactoryStorage(db, txID), usr, txID),
		SrvNursingConsultationRoutineReview:        nursing_consultation_routine_review.NewRoutineReviewService(nursing_consultation_routine_review.FactoryStorage(db, txID), usr, txID),
		SrvNursingConsultationAccompanyingData:     nursing_consultation_accompanying_data.NewAccompanyingDataService(nursing_consultation_accompanying_data.FactoryStorage(db, txID), usr, txID),
		SrvNursingConsultationVaccine:              nursing_consultation_vaccine.NewVaccineService(nursing_consultation_vaccine.FactoryStorage(db, txID), usr, txID),
		SrvNursingConsultationPhysicalTest:         nursing_consultation_physical_test.NewPhysicalTestService(nursing_consultation_physical_test.FactoryStorage(db, txID), usr, txID),
		SrvNursingConsultationLaboratoryTest:       nursing_consultation_laboratory_test.NewLaboratoryTestService(nursing_consultation_laboratory_test.FactoryStorage(db, txID), usr, txID),
		SrvNursingConsultationPreferentialTest:     nursing_consultation_preferential_test.NewPreferentialTestService(nursing_consultation_preferential_test.FactoryStorage(db, txID), usr, txID),
		SrvNursingConsultationSexualityTest:        nursing_consultation_sexuality_test.NewSexualityTestService(nursing_consultation_sexuality_test.FactoryStorage(db, txID), usr, txID),
		SrvNursingConsultationVisualTest:           nursing_consultation_visual_test.NewVisualTestService(nursing_consultation_visual_test.FactoryStorage(db, txID), usr, txID),
		SrvNursingConsultationMedicationTreatment:  nursing_consultation_medication_treatment.NewMedicationTreatmentService(nursing_consultation_medication_treatment.FactoryStorage(db, txID), usr, txID),
		SrvNursingConsultationPerformedProcedures:  nursing_consultation_performed_procedures.NewPerformedProceduresService(nursing_consultation_performed_procedures.FactoryStorage(db, txID), usr, txID),
		SrvDentistryConsultation:                   dentistry_consultation.NewDentistryConsultationService(dentistry_consultation.FactoryStorage(db, txID), usr, txID),
		SrvDentistryConsultationBuccalTest:         dentistry_consultation_buccal_test.NewBuccalTestService(dentistry_consultation_buccal_test.FactoryStorage(db, txID), usr, txID),
		SrvDentistryConsultationOdontogramReview:   dentistry_consultation_odontogram_review.NewOdontogramReviewService(dentistry_consultation_odontogram_review.FactoryStorage(db, txID), usr, txID),
		SrvDentistryConsultationBuccalConsultation: dentistry_consultation_buccal_consultation.NewBuccalConsultationService(dentistry_consultation_buccal_consultation.FactoryStorage(db, txID), usr, txID),
		SrvDentistryConsultationBuccalProcedure:    dentistry_consultation_buccal_procedure.NewBuccalProcedureService(dentistry_consultation_buccal_procedure.FactoryStorage(db, txID), usr, txID),
		SrvMedicalConsultation:                     medical_consultation.NewMedicalConsultationService(medical_consultation.FactoryStorage(db, txID), usr, txID),
		SrvConsultationIntegralAttention:           consultation_integral_attention.NewConsultationIntegralAttentionService(consultation_integral_attention.FactoryStorage(db, txID), usr, txID),
		SrvAnnouncementSignatures:                  announcement_signatures.NewAnnouncementSignaturesService(announcement_signatures.FactoryStorage(db, txID), usr, txID),
		SrvConsultationAssignment:                  consultation_assignment.NewConsultationAssignmentService(consultation_assignment.FactoryStorage(db, txID), usr, txID),
		SrvMedicalGeneralMedicineConsultation:      medical_general_medicine_consultation.NewGeneralMedicineConsultationService(medical_general_medicine_consultation.FactoryStorage(db, txID), usr, txID),
		SrvPaymentsConcept:                         paymentsconcept.NewPaymentConceptService(paymentsconcept.FactoryStorage(db, txID), usr, txID),
	}
}
