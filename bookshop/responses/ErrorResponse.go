package responses

import (
	"bookshop/interfaces"
	"net/http"
)

type Error struct {
	interfaces.Response
	Data string `json:"data"`
}

func ErrorResponse(w http.ResponseWriter, status int, err string, data string) {
	var resp Error

	resp.Status = status
	resp.Error = err
	resp.Data = data

	resp.Send(w, status, resp)
}
