package payments

import "github.com/asaskevich/govalidator"

type RequestPayment struct {
	TipoServicio   string `json:"tipo_servicio" valid:"required"`
	NombreServicio string `json:"nombre_servicio" valid:"required"`
	Dni            string `json:"dni" valid:"required"`
	Recibo         string `json:"recibo" valid:"required"`
}

func (m *RequestPayment) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
