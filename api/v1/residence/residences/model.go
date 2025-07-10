package residences

import "dbu-api/internal/models"

type ResponseResidence struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
}

type ResponseStudentsResidence struct {
	Students []*models.Student `json:"students"`
	Total    int               `json:"total"`
}
