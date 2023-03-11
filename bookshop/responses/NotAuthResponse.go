package responses

import (
	"net/http"
)

type NotAuth struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
	Data   string `json:"data"`
}

func NotAuthResponse(w http.ResponseWriter) {
	ErrorResponse(w, 401, "Not auth", "")
}
