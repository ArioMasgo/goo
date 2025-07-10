package residence

import (
	"dbu-api/internal/dbx"
	"dbu-api/internal/models"
	"dbu-api/pkg/orchestrator/low_code_residences"
	"github.com/google/uuid"
	"testing"
)

type ResidenceTestCase struct {
	name         string
	input        models.Residence
	expectedCode int
	expectError  bool
	expectedNil  bool
}

func TestCreatedResidences(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	user := models.User{
		ID: 1,
	}

	testCases := []ResidenceTestCase{
		{
			name: "Successful Create",
			input: models.Residence{
				ID:          uuid.New().String(),
				Name:        "Residencia Uno",
				Gender:      "femenino",
				Description: "Descripción de prueba",
				Address:     "Dirección de prueba",
				Status:      "habilitado",
				Floors: []*models.Floor{
					{
						Floor: 1,
						Rooms: []models.Room{
							{
								ID:       uuid.New().String(),
								Number:   1,
								Capacity: 2,
								Status:   "habilitado",
							},
							{
								ID:       uuid.New().String(),
								Number:   2,
								Capacity: 2,
								Status:   "habilitado",
							},
						},
					},
				},
			},
			expectedCode: 20,
			expectError:  false,
			expectedNil:  false,
		},
		{
			name: "Failed Create - Invalid Gender",
			input: models.Residence{
				ID:          uuid.New().String(),
				Name:        "Residencia dos",
				Gender:      "ds",
				Description: "Descripción de prueba",
				Address:     "Dirección de prueba",
				Status:      "habilitado",
				Floors: []*models.Floor{
					{
						Floor: 1,
						Rooms: []models.Room{
							{
								ID:       uuid.New().String(),
								Number:   1,
								Capacity: 2,
								Status:   "habilitado",
							},
							{
								ID:       uuid.New().String(),
								Number:   2,
								Capacity: 2,
								Status:   "habilitado",
							},
						},
					},
				},
			},
			expectedCode: 102,
			expectError:  true,
			expectedNil:  false,
		},
		{
			name: "Failed Create - Invalid Status",
			input: models.Residence{
				ID:          uuid.New().String(),
				Name:        "Residencia tres",
				Gender:      "femenino",
				Description: "Descripción de prueba",
				Address:     "Dirección de prueba",
				Status:      "kk",
				Floors: []*models.Floor{
					{
						Floor: 1,
						Rooms: []models.Room{
							{
								ID:       uuid.New().String(),
								Number:   1,
								Capacity: 2,
								Status:   "habilitado",
							},
							{
								ID:       uuid.New().String(),
								Number:   2,
								Capacity: 2,
								Status:   "habilitado",
							},
						},
					},
				},
			},
			expectedCode: 102,
			expectError:  true,
			expectedNil:  false,
		},
		{
			name: "Failed Create - without floors",
			input: models.Residence{
				ID:          uuid.New().String(),
				Name:        "Residencia cuatro",
				Gender:      "femenino",
				Description: "Descripción de prueba",
				Address:     "Dirección de prueba",
				Status:      "habilitado",
			},
			expectedCode: 101,
			expectError:  true,
			expectedNil:  false,
		},
		{
			name: "Failed Create - without rooms",
			input: models.Residence{
				ID:          uuid.New().String(),
				Name:        "Residencia cinco",
				Gender:      "femenino",
				Description: "Descripción de prueba",
				Address:     "Dirección de prueba",
				Status:      "habilitado",
				Floors: []*models.Floor{
					{
						Floor: 1,
					},
				},
			},
			expectedCode: 202,
			expectError:  true,
			expectedNil:  false,
		},
		{
			name: "Failed Create - residence register exists",
			input: models.Residence{
				ID:          uuid.New().String(),
				Name:        "Residencia Uno",
				Gender:      "femenino",
				Description: "Descripción de prueba",
				Address:     "Dirección de prueba",
				Status:      "habilitado",
				Floors: []*models.Floor{
					{
						Floor: 1,
						Rooms: []models.Room{
							{
								ID:       uuid.New().String(),
								Number:   1,
								Capacity: 2,
								Status:   "habilitado",
							},
						},
					},
				},
			},
			expectedCode: 201,
			expectError:  true,
			expectedNil:  false,
		},
	}

	srv := low_code_residences.NewResidence(db, &user, txID)

	for _, tc := range testCases {
		tc := tc // Capturar la variable tc para evitar problemas con goroutines
		t.Run(tc.name, func(t *testing.T) {
			// Llamar a la función UpdateResidence
			code, err := srv.CreateResidenceLowCode(&tc.input)

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
				if tc.expectedNil {
					t.Errorf("esperaba data no nula pero obtuvo nula")
				}
			}
		})
	}
}

func TestUpdatedResidences(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	user := models.User{
		ID: 1,
	}

	testCases := []ResidenceTestCase{
		{
			name: "Successful Updated",
			input: models.Residence{
				ID:          uuid.New().String(),
				Name:        "Residencia Uno",
				Gender:      "femenino",
				Description: "Descripción de prueba",
				Address:     "Dirección de prueba",
				Status:      "habilitado",
			},
			expectedCode: 20,
			expectError:  false,
			expectedNil:  false,
		},
		{
			name: "Failed Create - Invalid Gender",
			input: models.Residence{
				ID:          uuid.New().String(),
				Name:        "Residencia dos",
				Gender:      "ds",
				Description: "Descripción de prueba",
				Address:     "Dirección de prueba",
				Status:      "habilitado",
			},
			expectedCode: 102,
			expectError:  true,
			expectedNil:  false,
		},
		{
			name: "Failed Create - Invalid Status",
			input: models.Residence{
				ID:          uuid.New().String(),
				Name:        "Residencia tres",
				Gender:      "femenino",
				Description: "Descripción de prueba",
				Address:     "Dirección de prueba",
				Status:      "kk",
			},
			expectedCode: 102,
			expectError:  true,
			expectedNil:  false,
		},
	}

	srv := low_code_residences.NewResidence(db, &user, txID)

	for _, tc := range testCases {
		tc := tc // Capturar la variable tc para evitar problemas con goroutines
		t.Run(tc.name, func(t *testing.T) {
			// Llamar a la función UpdateResidence
			code, err := srv.CreateResidenceLowCode(&tc.input)

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
				if tc.expectedNil {
					t.Errorf("esperaba data no nula pero obtuvo nula")
				}
			}
		})
	}
}
