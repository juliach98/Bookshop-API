package responses

import (
	"net/http"
)

func NoPermsResponse(w http.ResponseWriter) {
	ErrorResponse(w, 403, "Not granted", "")
}
