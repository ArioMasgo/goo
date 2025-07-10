package ws

import (
	"bytes"
	"dbu-api/internal/logger"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func CallWebService(service WebServiceSchemeRequest) ([]byte, int, error) {
	reqURL, err := url.Parse(service.Url)
	if err != nil {
		return nil, 1, fmt.Errorf("no se pudo parsear la URL: %v", err)
	}

	// Agregar parámetros de consulta (query params) si existen
	if len(service.Params) > 0 {
		query := reqURL.Query()
		for key, value := range service.Params {
			query.Add(key, value)
		}
		reqURL.RawQuery = query.Encode()
		service.Url = reqURL.String()
	}

	var payload io.Reader
	// Solo agregar el cuerpo si no es un método GET
	if service.Method != http.MethodGet {
		payload = strings.NewReader(service.Payload)
	}

	request, err := http.NewRequest(service.Method, service.Url, payload)
	if err != nil {
		logger.Error.Printf("no se  puedo obtener respuesta: %v", err)
		return nil, 1, err
	}

	for key, value := range service.Headers {
		request.Header.Add(key, value)
	}

	client := &http.Client{}

	responseClient, err := client.Do(request)

	statusCode := 0

	if err != nil {
		logger.Error.Printf("no se  puedo enviar la petición: %v", err)
		return nil, statusCode, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error.Printf("no se pudo ejecutar defer body close: %v", err)
		}
	}(responseClient.Body)

	responseBody, err := io.ReadAll(responseClient.Body)
	if err != nil {
		logger.Error.Printf("no se  puedo obtener respuesta: %v", err)
		return responseBody, responseClient.StatusCode, err
	}

	return responseBody, responseClient.StatusCode, nil
}

func CallApiRest(method, url string, jsonBytes []byte, token string) (int, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		logger.Error.Printf("no se  puedo crear la petición: %v  -- log: ", err)
		return 500, nil, err
	}

	if token != "" {
		req.Header.Set("Authorization", token)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error.Printf("no se  puedo enviar la petición: %v  -- log: ", err)
		return 500, nil, err
	}

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Printf("no se  puedo obtener response body: %v  -- log: ", err)
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, response, nil

}
