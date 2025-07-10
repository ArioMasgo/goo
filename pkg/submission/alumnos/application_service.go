package alumnos

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
)

type PortsServerAlumnos interface {
	GetStudentsAcceptedBySubmission(id int64) ([]*Alumno, error)
	GetStudentsAcceptedBySubmissionNewbie(id int64) ([]*Alumno, error)
	GetStudentsByResidenceANDBySubmission(residenceID string, submissionID int64, page, limit int, filter string) ([]*StudentInformation, error)
	GetTotalStudentsByResidenceANDBySubmission(residenceID string, submissionID int64, filter string) (int, error)
	GetStudentsBySubmission(submissionID int, page, limit int, gender string, statusService string, departmentRequirementID int) ([]*StudentInformationSubmission, error)
	GetTotalStudentsBySubmission(submissionID int, gender string, statusService string, departmentRequirementID int) (int, error)
	GetStudentsBySubmissionExcel(submissionID int) ([]*models.StudentExcel, error)
	GetStudentsByRooms(rooms []string, submissionID int64) ([]*StudentInformation, error)
	GetStudentAcceptedBySubmission(submissionID, studentID int64) (*Alumno, int, error)
}

type service struct {
	repository ServicesAlumnosRepository
	user       *models.User
	txID       string
}

func NewAlumnosService(repository ServicesAlumnosRepository, user *models.User, TxID string) PortsServerAlumnos {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) GetStudentsAcceptedBySubmission(id int64) ([]*Alumno, error) {
	return s.repository.getStudentsAcceptedBySubmission(id)
}

func (s *service) GetStudentsAcceptedBySubmissionNewbie(id int64) ([]*Alumno, error) {
	return s.repository.getStudentsAcceptedBySubmissionNewbie(id)
}

func (s *service) GetStudentsByResidenceANDBySubmission(residenceID string, submissionID int64, page, limit int, filter string) ([]*StudentInformation, error) {
	return s.repository.getStudentsByResidenceANDBySubmission(residenceID, submissionID, page, limit, filter)
}

func (s *service) GetTotalStudentsByResidenceANDBySubmission(residenceID string, submissionID int64, filter string) (int, error) {
	return s.repository.getTotalStudentsByResidenceANDBySubmission(residenceID, submissionID, filter)
}

func (s *service) GetStudentsBySubmission(submissionID int, page, limit int, gender string, statusService string, departmentRequirementID int) ([]*StudentInformationSubmission, error) {
	return s.repository.getStudentsBySubmission(submissionID, page, limit, gender, statusService, departmentRequirementID)
}

func (s *service) GetTotalStudentsBySubmission(submissionID int, gender string, statusService string, departmentRequirementID int) (int, error) {
	return s.repository.getTotalStudentsBySubmission(submissionID, gender, statusService, departmentRequirementID)
}

func (s *service) GetStudentsBySubmissionExcel(submissionID int) ([]*models.StudentExcel, error) {
	return s.repository.getStudentsBySubmissionExcel(submissionID)
}

func (s *service) GetStudentsByRooms(rooms []string, submissionID int64) ([]*StudentInformation, error) {
	return s.repository.getStudentsByRooms(rooms, submissionID)
}

func (s *service) GetStudentAcceptedBySubmission(submissionID, studentID int64) (*Alumno, int, error) {
	if submissionID < 1 || studentID < 1 {
		logger.Error.Println(s.txID, " - couldn't meet validations:")
		return nil, 15, fmt.Errorf("couldn't meet validations")
	}
	m, err := s.repository.getStudentAcceptedBySubmission(submissionID, studentID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}
