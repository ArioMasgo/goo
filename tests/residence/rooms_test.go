package residence

import (
	"dbu-api/internal/dbx"
	"dbu-api/internal/models"
	"dbu-api/pkg/residence"
	"dbu-api/pkg/residence/rooms"
	"github.com/google/uuid"
	"testing"
)

type RoomStructTest struct {
	ID          string `db:"id" json:"id" valid:"uuid,required"`
	Number      int    `db:"numero" json:"number" valid:"required"`
	Capacity    int    `db:"capacidad" json:"capacity" valid:"required"`
	Status      string `db:"estado" json:"status" valid:"required"`
	Floor       int    `db:"piso" json:"floor" valid:"required"`
	ResidenceID string `db:"residencia_id" json:"residence_id" valid:"required"`
}

type RoomTestCase struct {
	name         string
	input        rooms.Room
	expectedCode int
	expectError  bool
	expectedNil  bool
}

func TestCreatedRoom(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	user := models.User{
		ID: 1,
	}

	srv := residence.NewServerResidence(db, &user, txID)

	testCases := []RoomTestCase{
		{
			name: "Successful Create",
			input: rooms.Room{
				ID:          uuid.New().String(),
				Number:      2,
				Capacity:    3,
				Status:      "habilitado",
				Floor:       1,
				ResidenceID: "4bb4b3ac-2616-47aa-95f8-0afe200fd9cc",
			},
			expectedCode: 29,
			expectError:  false,
			expectedNil:  false,
		},
		{
			name: "Failed Create - Missing Status",
			input: rooms.Room{
				ID:          uuid.New().String(),
				Number:      2,
				Capacity:    3,
				Status:      "",
				Floor:       1,
				ResidenceID: "4bb4b3ac-2616-47aa-95f8-0afe200fd9cc",
			},
			expectedCode: 15,
			expectError:  false,
			expectedNil:  false,
		},
		{
			name: "Failed Create - Invalid UUID",
			input: rooms.Room{
				ID:          "invalid",
				Number:      2,
				Capacity:    3,
				Status:      "",
				Floor:       1,
				ResidenceID: "4bb4b3ac-2616-47aa-95f8-0afe200fd9cc",
			},
			expectedCode: 15,
			expectError:  false,
			expectedNil:  false,
		},
		{
			name: "Failed Create - Duplicity number",
			input: rooms.Room{
				ID:          "invalid",
				Number:      2,
				Capacity:    3,
				Status:      "",
				Floor:       1,
				ResidenceID: "4bb4b3ac-2616-47aa-95f8-0afe200fd9cc",
			},
			expectedCode: 15,
			expectError:  false,
			expectedNil:  false,
		},
	}

	for _, tc := range testCases {
		tc := tc // Capturar la variable tc para evitar problemas con goroutines
		t.Run(tc.name, func(t *testing.T) {
			// Llamar a la función CreateRoom
			data, code, err := srv.SrvRoom.CreateRoom(
				tc.input.ID,
				tc.input.Number,
				tc.input.ResidenceID,
				tc.input.Capacity,
				tc.input.Status,
				tc.input.Floor,
			)

			// Verificar si se espera un error
			if tc.expectError {
				if err == nil {
					t.Errorf("esperaba error pero no lo obtuvo")
				}
				if code != tc.expectedCode {
					t.Errorf("esperaba código %d pero obtuvo %d", tc.expectedCode, code)
				}
			} else {
				if err != nil {
					t.Errorf("no esperaba error pero obtuvo: %v", err)
				}
				if code != tc.expectedCode {
					t.Errorf("esperaba código %d pero obtuvo %d", tc.expectedCode, code)
				}
				if tc.expectedNil && data == nil {
					t.Errorf("esperaba data no nula pero obtuvo nula")
				}
			}

			// Verificar si se espera que data sea nil
			if tc.expectedNil && data != nil {
				t.Errorf("esperaba data nil pero obtuvo: %v", data)
			}
		})
	}
}
