package users

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.LoginUser

	err := req.Load(r)

	if err != nil {
		l.Print("User_controller", "error", "Login", "", "req.Load", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	if len(req.Email) == 0 {
		l.Print("User_controller", "error", "Login", "", "len(req.Email) == 0", "Email is empty")
		responses.ErrorResponse(w, 502, "Email is empty", "")
		return
	}

	if len(req.Password) == 0 {
		l.Print("User_controller", "error", "Login", "", "len(req.Password) == 0", "Password is empty")
		responses.ErrorResponse(w, 503, "Password is empty", "")
		return
	}

	var user models.User

	user.Email = req.Email
	user.Password = helpers.GetEncryptedPassword(req.Password)

	found, err := user.Login()

	if err != nil {
		l.Print("User_controller", "error", "Login", "", "user.Login", err.Error())
		responses.ErrorResponse(w, 504, err.Error(), "")
		return
	}

	if !found {
		l.Print("User_controller", "error", "Login", "", "user.Login", "User not found")
		responses.ErrorResponse(w, 404, "User not found", "")
		return
	}

	var userToken models.UserToken
	userToken.UserID = user.ID

	err = userToken.Delete()

	if err != nil {
		l.Print("User_controller", "error", "Login", "", "userToken.Delete", err.Error())
		responses.ErrorResponse(w, 505, err.Error(), "")
		return
	}

	err = userToken.CreateToken(user.ID)
	if err != nil {
		l.Print("User_controller", "error", "Login", "", "userToken.CreateToken", err.Error())
		responses.ErrorResponse(w, 506, err.Error(), "")
		return
	}

	var resp responses.RegisterUser
	resp.Fill(userToken.Token)
	resp.Send(w, 200, resp)
}
