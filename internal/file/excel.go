package file

import (
	"dbu-api/internal/models"
	"encoding/base64"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CreateExcelFile(m *models.ExcelFile) (string, int) {
	// Crear nuevo archivo Excel
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// Crear estilos
	headerStyle, err := createHeaderStyle(f)
	if err != nil {
		return "", 29
	}

	dataStyle, err := createDataStyle(f)
	if err != nil {
		return "", 29
	}

	// Procesar cada página
	for pageIndex, page := range m.Page {
		// Crear nueva hoja
		sheetName := page.Name
		if sheetName == "" {
			sheetName = fmt.Sprintf("Sheet%d", pageIndex+1)
		}

		index, _ := f.NewSheet(sheetName)

		// Procesar filas
		for i, row := range page.Rows {
			// Procesar columnas de cada fila
			for _, col := range row.Columns {
				cellName, err := excelize.JoinCellName(col.Column, row.Row)
				if err != nil {
					return "", 28
				}

				if err := f.SetCellValue(sheetName, cellName, col.Value); err != nil {
					return "", 27
				}

				// Aplicar estilo según si es encabezado o datos
				if i == 0 {
					if err := f.SetCellStyle(sheetName, cellName, cellName, headerStyle); err != nil {
						return "", 30
					}
				} else {
					if err := f.SetCellStyle(sheetName, cellName, cellName, dataStyle); err != nil {
						return "", 30
					}
				}
			}
		}

		// Aplicar autoajuste y otros estilos
		if err := applySheetStyle(f, sheetName); err != nil {
			return "", 31
		}

		// Establecer como hoja activa si es la primera
		if pageIndex == 0 {
			f.SetActiveSheet(index)
		}
	}

	// Eliminar la hoja por defecto (Sheet1)
	if len(m.Page) > 0 {
		f.DeleteSheet("Sheet1")
	}

	// Generar nombre de archivo
	if m.Name == "" {
		m.Name = fmt.Sprintf("report_%s.xlsx", time.Now().Format("20060102_150405"))
	}

	if !strings.HasSuffix(m.Name, ".xlsx") {
		m.Name += ".xlsx"
	}

	// Crear directorio
	outputDir := "./temp/excel"
	if m.Path != "" {
		outputDir = m.Path
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", 33
	}

	filePath := filepath.Join(outputDir, m.Name)

	if err := f.SaveAs(filePath); err != nil {
		return "", 18
	}

	buffer, err := f.WriteToBuffer()
	if err != nil {
		return "", 18
	}

	base64String := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return base64String, 223
}

func createHeaderStyle(f *excelize.File) (int, error) {
	return f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Size:  11,
			Color: "#FFFFFF",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#34b489"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#005151", Style: 1},
			{Type: "top", Color: "#005151", Style: 1},
			{Type: "bottom", Color: "#005151", Style: 1},
			{Type: "right", Color: "#005151", Style: 1},
		},
	})
}

func createDataStyle(f *excelize.File) (int, error) {
	return f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size: 10,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#005151", Style: 1},
			{Type: "top", Color: "#005151", Style: 1},
			{Type: "bottom", Color: "#005151", Style: 1},
			{Type: "right", Color: "#005151", Style: 1},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFFF"},
			Pattern: 1,
		},
	})
}

// applySheetStyle aplica estilos básicos a la hoja
func applySheetStyle(f *excelize.File, sheetName string) error {
	cols, err := f.GetCols(sheetName)
	if err != nil {
		return err
	}

	for idx, col := range cols {
		maxWidth := 0
		for _, cell := range col {
			cellWidth := len(cell) + 2 // +2 para un poco de padding
			if cellWidth > maxWidth {
				maxWidth = cellWidth
			}
		}

		// Limitar el ancho máximo
		if maxWidth > 50 {
			maxWidth = 50
		}

		colName, err := excelize.ColumnNumberToName(idx + 1)
		if err != nil {
			return err
		}

		if err := f.SetColWidth(sheetName, colName, colName, float64(maxWidth)); err != nil {
			return err
		}
	}

	return nil
}
