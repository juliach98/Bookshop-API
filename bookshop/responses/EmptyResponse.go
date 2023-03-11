package responses

import (
	"bookshop/interfaces"
	"net/http"
)

type emptyR struct {
	interfaces.Response
	Data string `json:"data"`
}

func EmptyResponse(w http.ResponseWriter) {
	var resp emptyR

	resp.Status = 200
	resp.Error = ""
	resp.Data = ""

	resp.Send(w, 200, resp)
}
