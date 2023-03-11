package middlewares

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/responses"
	"net/http"
)

func IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		var l helpers.Logger

		if r.Context().Value("user").(models.User).IsAdmin != true {
			responses.NoPermsResponse(writer)
			l.Print("request", "info", r.URL.String(), "", "", "NotAuthResponse (Not admin)")
			return
		}

		next.ServeHTTP(writer, r)
	}
}
