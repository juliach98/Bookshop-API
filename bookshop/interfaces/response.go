package interfaces

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func (m *Response) Send(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if status != 200 {
		w.WriteHeader(status)
	}
	responseByte, _ := json.Marshal(data)
	_, _ = w.Write(responseByte)
}
