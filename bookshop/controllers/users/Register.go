package users

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.RegisterUser

	err := req.Load(r)

	if err != nil {
		l.Print("User_controller", "error", "Register", "", "req.Load", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	if len(req.Name) == 0 {
		l.Print("User_controller", "error", "Register", "", "len(req.Name) == 0", "Name is empty")
		responses.ErrorResponse(w, 502, "Name is empty", "")
		return
	}

	if len(req.Surname) == 0 {
		l.Print("User_controller", "error", "Register", "", "len(req.Surname) == 0", "Surname is empty")
		responses.ErrorResponse(w, 503, "Surname is empty", "")
		return
	}

	if len(req.Email) == 0 {
		l.Print("User_controller", "error", "Register", "", "len(req.Email) == 0", "Email is empty")
		responses.ErrorResponse(w, 504, "Email is empty", "")
		return
	}

	if len(req.PhoneNumber) == 0 {
		l.Print("User_controller", "error", "Register", "", "len(req.PhoneNumber) == 0", "PhoneNumber is empty")
		responses.ErrorResponse(w, 505, "PhoneNumber is empty", "")
		return
	}

	if len(req.Password) == 0 {
		l.Print("User_controller", "error", "Register", "", "len(req.Password) == 0", "Password is empty")
		responses.ErrorResponse(w, 506, "Password is empty", "")
		return
	}

	var user models.User

	user.Name = req.Name
	user.Surname = req.Surname
	user.Email = req.Email
	user.PhoneNumber = req.PhoneNumber
	user.Password = helpers.GetEncryptedPassword(req.Password)

	found, err := user.FindByEmail()

	if err != nil {
		l.Print("User_controller", "error", "Register", "", "user.FindByEmail", err.Error())
		responses.ErrorResponse(w, 507, err.Error(), "")
		return
	}

	if found {
		l.Print("User_controller", "error", "Register", "", "users.FindByEmail", "User with this email is already registered")
		responses.ErrorResponse(w, 508, "User with this email is already registered", "")
		return
	}

	err = user.Create()

	if err != nil {
		l.Print("User_controller", "error", "Register", "", "user.Create", err.Error())
		responses.ErrorResponse(w, 509, err.Error(), "")
		return
	}

	var userToken models.UserToken

	err = userToken.CreateToken(user.ID)
	if err != nil {
		l.Print("User_controller", "error", "Register", "", "userToken.CreateToken", err.Error())
		responses.ErrorResponse(w, 510, err.Error(), "")
		return
	}

	var resp responses.RegisterUser
	resp.Fill(userToken.Token)
	resp.Send(w, 200, resp)
}
