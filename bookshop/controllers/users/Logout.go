package users

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/responses"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger

	user := r.Context().Value("user").(models.User)

	var token models.UserToken
	token.UserID = user.ID

	err := token.Delete()

	if err != nil {
		l.Print("User_controller", "error", "Logout", "", "token.Delete", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	responses.EmptyResponse(w)
}
