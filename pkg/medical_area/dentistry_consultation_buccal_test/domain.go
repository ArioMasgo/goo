package dentistry_consultation_buccal_test

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type BuccalTest struct {
	ID                    string `json:"id" db:"id"`
	IDConsultaOdontologia string `json:"consulta_odontologia_id" db:"consulta_odontologia_id"`
	OdontogramaIMG        string `json:"odontograma_img" db:"odontograma_img" valid:"-"`
	CPOD                  string `json:"cpod" db:"cpod" valid:"-"`
	Observacion           string `json:"observacion" db:"observacion" valid:"-"`
	IHOS                  string `json:"ihos" db:"ihos" valid:"-"`
	Comentarios           string `json:"comentarios" db:"comentarios" valid:"-"`

	IsDeleted   bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator string     `json:"user_creator" db:"user_creator"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

func NewBuccalTest(id string, consulta_odontologia_id string, odontograma_img string, cpod string, observacion string, ihos string, comentarios string) *BuccalTest {
	now := time.Now()
	return &BuccalTest{
		ID:                    id,
		IDConsultaOdontologia: consulta_odontologia_id,
		OdontogramaIMG:        odontograma_img,
		CPOD:                  cpod,
		Observacion:           observacion,
		IHOS:                  ihos,
		Comentarios:           comentarios,
		IsDeleted:             false,
		CreatedAt:             &now,
		UpdatedAt:             &now,
	}
}

func (m *BuccalTest) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
