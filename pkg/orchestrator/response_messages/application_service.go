package response_messages

import (
	"dbu-api/internal/logger"
	"fmt"
)

type Service struct {
	repository ServicesMessageRepository
}

func NewMessageService(repository ServicesMessageRepository) Service {
	return Service{repository: repository}
}

func (s Service) GetMessageByID(id int) (*ResponseMessage, int, error) {
	if id <= 0 {
		logger.Error.Println(" - don't meet validations:", fmt.Errorf("id isn't int"))
		return nil, 15, fmt.Errorf("id isn't int")
	}
	m, err := s.repository.GetByID(id)
	if err != nil {
		logger.Error.Println(" - couldn't getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}
