package models

type Metadata struct {
	Links      Links `json:"links"`
	Total      int64 `json:"total"`
	Offset     int64 `json:"offset"`
	Limit      int64 `json:"limit"`
	Page       int64 `json:"page"`
	TotalPages int64 `json:"total_pages"`
}

type Links struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
