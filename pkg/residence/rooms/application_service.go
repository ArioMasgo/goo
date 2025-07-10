package rooms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"

	"dbu-api/internal/logger"
	"dbu-api/internal/models"
)

type PortsServerRoom interface {
	CreateRoom(id string, number int, residenceID string, capacity int, status string, floor int) (*Room, int, error)
	UpdateRoom(id string, number int, residenceID string, capacity int, status string, floor int) (*Room, int, error)
	DeleteRoom(id string) (int, error)
	GetRoomByID(id string) (*Room, int, error)
	GetAllRooms() ([]*Room, error)
	GetAllRoomsByResidenceID(id string) ([]*Room, error)
	MultiCreate(residenceID string, m []*models.Floor) ([]*Room, error)
	UpdateOnlyCharacteristicsRoom(id string, capacity int, status string) (*Room, int, error)
	GetAllRoomsBySubmissionIDResidenceID(submissionID int64, residenceID string) ([]*Room, error)
	GetRoomsByResidence(residenceID string, page, limit int) ([]*Room, error)
}

type service struct {
	repository ServicesRoomRepository
	user       *models.User
	txID       string
}

func NewRoomService(repository ServicesRoomRepository, user *models.User, TxID string) PortsServerRoom {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateRoom(id string, number int, residenceID string, capacity int, status string, floor int) (*Room, int, error) {
	m := NewRoom(id, number, residenceID, capacity, status, floor, s.user.ID)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Room :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateRoom(id string, number int, residenceID string, capacity int, status string, floor int) (*Room, int, error) {
	m := NewRoom(id, number, residenceID, capacity, status, floor, s.user.ID)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Room :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteRoom(id string) (int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return 15, fmt.Errorf("id is required")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "rows affected error" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetRoomByID(id string) (*Room, int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}

	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllRooms() ([]*Room, error) {
	return s.repository.getAll()
}

func (s *service) GetAllRoomsByResidenceID(id string) ([]*Room, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - not valid uuid:", fmt.Errorf("id is required"))
		return nil, fmt.Errorf("id is required")
	}
	return s.repository.getAllRoomsByResidenceID(id)
}

func (s *service) MultiCreate(residenceID string, m []*models.Floor) ([]*Room, error) {
	var rooms []*Room
	for _, floor := range m {
		for _, room := range floor.Rooms {
			newRoom := NewRoom(room.ID, room.Number, residenceID, room.Capacity, room.Status, floor.Floor, s.user.ID)
			rooms = append(rooms, newRoom)
		}
	}

	if rooms == nil {
		return nil, fmt.Errorf("no exists rooms")
	}

	if len(rooms) == 0 {
		return nil, fmt.Errorf("no rooms found")
	}

	err := s.repository.multiCreate(rooms)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, err
	}

	return rooms, nil
}

func (s *service) UpdateOnlyCharacteristicsRoom(id string, capacity int, status string) (*Room, int, error) {
	m := NewRoom(id, 0, "", capacity, status, 0, s.user.ID)
	if err := s.repository.updateOnlyCharacteristics(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Room :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) GetAllRoomsBySubmissionIDResidenceID(submissionID int64, residenceID string) ([]*Room, error) {
	if !govalidator.IsUUID(residenceID) {
		logger.Error.Println(s.txID, " - not valid uuid:", fmt.Errorf("id is required"))
		return nil, fmt.Errorf("id is required")
	}
	return s.repository.gtAllRoomsBySubmissionIDResidenceID(submissionID, residenceID)
}

func (s *service) GetRoomsByResidence(residenceID string, page, limit int) ([]*Room, error) {
	if !govalidator.IsUUID(residenceID) {
		logger.Error.Println(s.txID, " - not valid uuid:", fmt.Errorf("id is required"))
		return nil, fmt.Errorf("id is required")
	}

	return s.repository.getRoomsByResidence(residenceID, page, limit)
}
