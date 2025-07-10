package models

type Response struct {
	Error    bool      `json:"error"`
	Msg      string    `json:"msg"`
	Data     any       `json:"data"`
	Code     int       `json:"code"`
	Type     string    `json:"type"`
	Metadata *Metadata `json:"metadata"`
}
