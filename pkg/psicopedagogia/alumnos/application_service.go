package alumnos

import (
	"fmt"
)

type PortsServerEstudiante interface {
	GetEstudianteByDni(dni string, tipoParticipante string) (any, int, error)
	GetNameEstudianteByDni(dni string) (*BasicEstudiante, int, error)
}

type service struct {
	repository ServicesEstudianteRepository
	txID       string
}

func NewEstudianteService(repository ServicesEstudianteRepository, txID string) PortsServerEstudiante {
	return &service{repository: repository, txID: txID}
}

func (s *service) GetEstudianteByDni(dni string, tipoParticipante string) (any, int, error) {
	if dni == "" {
		return nil, 400, fmt.Errorf("DNI is required")
	}

	participante, err2 := s.repository.GetParticipanteByDNI(dni, tipoParticipante)
	if err2 == nil && participante != nil {
		return participante, 200, nil
	}

	if tipoParticipante == "Alumno" {
		student, err := s.repository.getByDni(dni)
		if err == nil && student != nil {
			return student, 200, nil
		}
	}

	return nil, 404, fmt.Errorf("no se encontró un estudiante o participante con el DNI proporcionado")
}

func (s *service) GetNameEstudianteByDni(dni string) (*BasicEstudiante, int, error) {
	if dni == "" {
		return nil, 400, fmt.Errorf("DNI is required")
	}

	stundent, err := s.repository.getNameByDni(dni)
	if err != nil {
		return nil, 404, err
	}

	return stundent, 200, nil
}
