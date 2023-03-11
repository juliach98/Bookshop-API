package users

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.RegisterUser

	err := req.Load(r)

	if err != nil {
		l.Print("User_controller", "error", "Update", "", "req.Load", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	user := r.Context().Value("user").(models.User)

	if len(req.Name) > 0 {
		user.Name = req.Name
	}

	if len(req.Surname) > 0 {
		user.Surname = req.Surname
	}

	if len(req.Email) > 0 {
		user.Email = req.Email
	}

	if len(req.PhoneNumber) > 0 {
		user.PhoneNumber = req.PhoneNumber
	}

	err = user.Update()

	if err != nil {
		l.Print("User_controller", "error", "Update", "", "user.Update", err.Error())
		responses.ErrorResponse(w, 502, err.Error(), "")
		return
	}

	responses.EmptyResponse(w)
}
