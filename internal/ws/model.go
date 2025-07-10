package ws

import "time"

// Response estructura para respuesta estandar de api
type Response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
}

type FormDataRequest struct {
	Type  string `json:"type"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type FormUrlEncodedRequest struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type WebServiceSchemeRequest struct {
	Payload  string            `json:"payload"`
	Url      string            `json:"url"`
	Duration time.Duration     `json:"duration"`
	Method   string            `json:"method"`
	Headers  map[string]string `json:"headers"`
	Params   map[string]string `json:"params"`
}
