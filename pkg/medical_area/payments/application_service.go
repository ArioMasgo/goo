package payments_concept

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"dbu-api/internal/ws"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type PortsServerPaymentConcept interface {
	Search(dni, area, tipoServicio, servicio, recibo string) (*PagoTesoreria, int, error)
}

type service struct {
	repository ServicesPaymentConceptRepository
	user       *models.User
	txID       string
}

func NewPaymentConceptService(repository ServicesPaymentConceptRepository, user *models.User, TxID string) PortsServerPaymentConcept {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) Search(dni, area, tipoServicio, nombreServicio, recibo string) (*PagoTesoreria, int, error) {
	paymentConcept, err := s.repository.search(area, tipoServicio, nombreServicio)
	if err != nil {
		logger.Error.Println(s.txID, " - error searching payment concept:", err)
		return nil, 3, fmt.Errorf("error searching payment concept")
	}

	if paymentConcept == nil {
		logger.Error.Println(s.txID, " - not exist payment concept:", err)
		return nil, 255, fmt.Errorf("not exist payment concept")
	}

	if paymentConcept.RequierePago == false {
		return nil, 256, nil
	}

	code, response, err := ws.CallApiRest("GET", "https://tesoreria.unas.edu.pe/api/cajapaymentsunasdbu-normal-payments/"+dni, nil, "")
	if err != nil {
		logger.Error.Println(s.txID, " - error searching payment:", err)
		return nil, 251, fmt.Errorf("error searching payment")
	}

	if code != 200 {
		logger.Error.Println(s.txID, " - status code http = ", code, err)
		return nil, 251, fmt.Errorf("Ocurrió un error al buscar el pago")
	}
	detailPayment := DetallePagoTesoreria{}
	err = json.Unmarshal(response, &detailPayment)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't decode response", err)
		return nil, 251, fmt.Errorf("couldn't decode response")
	}

	payment, code := s.getPaymentByReceipt(recibo, paymentConcept.CodigoConcepto, detailPayment)
	if payment == nil {
		logger.Error.Println(s.txID, " - not exist payment", err)
		return nil, code, fmt.Errorf("not exist payment")
	}

	code, err = s.validatePayments(recibo, area, tipoServicio, nombreServicio, payment)
	if err != nil {
		logger.Error.Println(s.txID, " - not exist payment", err)
		return nil, code, err
	}

	return payment, 254, nil
}

func (s *service) getPaymentByReceipt(recibo string, codigoConcepto *int, detailPayment DetallePagoTesoreria) (*PagoTesoreria, int) {
	existReceipt := validateExistReceipt(recibo, detailPayment)
	if existReceipt == false {
		return nil, 250
	}

	for _, pago := range detailPayment.PagosRealizados {
		if (pago.CodRecibo != nil && *pago.CodRecibo == recibo) && pago.CodigoConcepto == *codigoConcepto {
			return &pago, 254
		}

		if (pago.CodReciboCanje != nil && *pago.CodReciboCanje == recibo) && pago.CodigoConcepto == *codigoConcepto {
			return &pago, 254
		}
	}

	return nil, 253
}

func validateExistReceipt(recibo string, detailPayment DetallePagoTesoreria) bool {
	for _, pago := range detailPayment.PagosRealizados {
		if pago.CodRecibo != nil && *pago.CodRecibo == recibo {
			return true
		}

		if pago.CodReciboCanje != nil && *pago.CodReciboCanje == recibo {
			return true
		}
	}

	return false
}

func (s *service) validatePayments(recibo, area, tipoServicio, nombreServicio string, payment *PagoTesoreria) (int, error) {
	paymentsArea, err := s.searchPaymentAreaService(recibo, area, tipoServicio, nombreServicio)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't search payment odontologia procedure", err)
		return 3, fmt.Errorf("couldn't search payment odontologia procedure")
	}

	cantidadStr := strings.TrimSuffix(payment.Cantidad, ".00")
	cantidad, err := strconv.Atoi(cantidadStr)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't convert amount", err)
		return 8, fmt.Errorf("couldn't convert amount")
	}

	if len(paymentsArea) >= cantidad {
		logger.Error.Println(s.txID, " - receipt has already been used", err)
		return 252, fmt.Errorf("receipt has already been used")
	}

	return 254, nil
}

func (s *service) searchPaymentAreaService(recibo, area, tipoServicio, servicio string) ([]*PagosServicios, error) {
	if area == "odontologia" {
		if tipoServicio == "procedimiento" {
			paymentServices, err := s.repository.searchPaymentProcedureOdontologia(recibo, servicio)
			if err != nil {
				return nil, err
			}

			return paymentServices, nil
		}
	}

	return nil, fmt.Errorf("Ocurrió un error al buscar el tipo de servicio")
}
