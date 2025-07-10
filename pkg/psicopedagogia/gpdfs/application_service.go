package gpdfs

import (
	"dbu-api/internal/file"
	"errors"
)

type PortsServerPDFs interface {
	GeneratePDF_SRQ(participantID int) ([]byte, error)
}

type service struct {
	repository ServicesGpdfsRepository
	txID       string
}

func NewPDFService(repository ServicesGpdfsRepository, txID string) PortsServerPDFs {
	return &service{repository: repository, txID: txID}
}

var (
	ErrInvalidParticipantID    = errors.New("invalid participant ID")
	ErrFetchingParticipantData = errors.New("error fetching participant data")
	ErrFetchingSurveyResponses = errors.New("error fetching survey responses")
	ErrGeneratingPDF           = errors.New("error generating PDF")
)

func (s *service) GeneratePDF_SRQ(participantID int) ([]byte, error) {
	if participantID == 0 {
		return nil, errors.New("invalid participant ID")
	}

	participant, err := s.repository.GetParticipantByID(participantID)
	if err != nil {
		return nil, errors.New("error fetching participant data")
	}

	responses, err := s.repository.GetResponsesPerParticipant(participantID, 1)
	if err != nil {
		return nil, errors.New("error fetching survey responses")
	}

	var responsesAny []any
	for _, r := range responses {
		responsesAny = append(responsesAny, r)
	}
	dataPdf, err := file.CreatePDF_SRQ(participant, responsesAny)
	if err != nil {
		return nil, errors.New("error generating PDF")
	}

	return dataPdf, nil
}
