package submissions

import "dbu-api/internal/models"

type ResponseStudentsSubmission struct {
	Students []*models.Student `json:"students"`
	Total    int               `json:"total"`
}
