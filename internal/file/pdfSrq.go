package file

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func CreatePDF_SRQ(participant any, responses []any) ([]byte, error) {
	htmlFile, err := os.ReadFile("internal/file/template/templateSRQResult.html")
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo HTML: %w", err)
	}

	tmpl, err := template.New("pdfTemplate").Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"cleanText": func(i interface{}) string {
			if str, ok := i.(string); ok {
				return strings.ReplaceAll(strings.ReplaceAll(str, "{", ""), "}", "")
			}
			return fmt.Sprintf("%v", i)
		},
	}).Parse(string(htmlFile))

	if err != nil {
		return nil, fmt.Errorf("error al parsear la plantilla HTML: %w", err)
	}

	var htmlBuffer bytes.Buffer
	err = tmpl.Execute(&htmlBuffer, struct {
		Participant any
		Questions   []any
	}{
		Participant: participant,
		Questions:   responses,
	})
	if err != nil {
		return nil, fmt.Errorf("error al renderizar la plantilla HTML: %w", err)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, fmt.Errorf("error al crear el generador PDF: %w", err)
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader(htmlBuffer.Bytes()))
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return nil, fmt.Errorf("error al generar el PDF: %w", err)
	}

	return pdfg.Bytes(), nil
}
