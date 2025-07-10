package announcement

import (
	"dbu-api/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAnnouncements(token string) ([]models.Announcement, error) {
	const url = "https://obu-dev.com/backend/convocatoria"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creando la solicitud: %w", err)
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en la solicitud: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo la respuesta: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error HTTP %d: %s", resp.StatusCode, string(body))
	}

	var announcements []models.Announcement
	err = json.Unmarshal(body, &announcements)
	if err != nil {
		return nil, fmt.Errorf("error parseando JSON: %w", err)
	}

	return announcements, err
}
