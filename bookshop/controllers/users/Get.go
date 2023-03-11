package users

import (
	"bookshop/models"
	"bookshop/responses"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)

	var resp responses.UserGet
	resp.Fill(user)
	resp.Send(w, 200, resp)
}
