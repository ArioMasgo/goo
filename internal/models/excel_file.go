package models

type ExcelFile struct {
	Name string      `json:"name"`
	Path string      `json:"path"`
	Page []ExcelPage `json:"page"`
}

type ExcelPage struct {
	Name string         `json:"name"`
	Rows []ExcelPageRow `json:"rows"`
}

type ExcelPageRow struct {
	Row     int               `json:"row"`
	Columns []ExcelPageColumn `json:"columns"`
}

type ExcelPageColumn struct {
	Column string `json:"column"`
	Value  string `json:"value"`
}

type MergeRange struct {
	Column1 string
	Column2 string
}
