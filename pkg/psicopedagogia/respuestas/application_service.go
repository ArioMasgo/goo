package respuestas

import (
	"fmt"
)

type PortsServerRespuestas interface {
	GetAll(page, pageSize int) ([]Respuesta, error)
	GetByID(id int) (*Respuesta, error)
	Create(res *Respuesta) error
	Update(id int, res *Respuesta) error
	Delete(id int) error
	GetResponsesPerParticipant(idParticipante int) ([]RespuestaDetalle, error)
	GetAllByParticipanteIdAndNumeroAtencion(idParticipante, numeroAtencion, page, pageSize int) ([]RespuestaDetalle, error)
}

type service struct {
	repository ServicesRespuestaRepository
	txID       string
}

func NewRespuestaService(repository ServicesRespuestaRepository, txID string) PortsServerRespuestas {
	return &service{repository: repository, txID: txID}
}

func (s *service) GetAll(page, pageSize int) ([]Respuesta, error) {
	return s.repository.GetAll(page, pageSize)
}

func (s *service) GetByID(id int) (*Respuesta, error) {
	return s.repository.GetByID(id)
}

func (s *service) Create(res *Respuesta) error {
	if res.IDParticipante == 0 || res.IDEncuesta == 0 || res.IDPregunta == 0 || res.Respuesta == "" {
		return fmt.Errorf("id_participante, id_encuesta, id_pregunta y respuesta son obligatorios")
	}
	return s.repository.Create(res)
}

func (s *service) Update(id int, res *Respuesta) error {
	if id == 0 {
		return fmt.Errorf("id es requerido")
	}
	return s.repository.Update(id, res)
}

func (s *service) Delete(id int) error {
	if id == 0 {
		return fmt.Errorf("id es requerido")
	}
	return s.repository.Delete(id)
}

func (s *service) GetResponsesPerParticipant(idParticipante int) ([]RespuestaDetalle, error) {
	if idParticipante == 0 {
		return nil, fmt.Errorf("id es requerido")
	}
	return s.repository.GetResponsesPerParticipant(idParticipante)
}

func (s *service) GetAllByParticipanteIdAndNumeroAtencion(idParticipante, numeroAtencion, page, pageSize int) ([]RespuestaDetalle, error) {
	if idParticipante == 0 {
		return nil, fmt.Errorf("idParticipante es requerido")
	}
	if numeroAtencion == 0 {
		return nil, fmt.Errorf("numeroAtencion es requerido")
	}
	respuestas, err := s.repository.GetAllByParticipanteIdAndNumeroAtencion(idParticipante, numeroAtencion, page, pageSize)
	if err != nil {
		return nil, err
	}
	return respuestas, nil
}
