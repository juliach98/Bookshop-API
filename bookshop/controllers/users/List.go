package users

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/responses"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger

	var user models.User

	userList, err := user.List()
	if err != nil {
		l.Print("User_controller", "error", "List", "", "user.List", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	var resp responses.UserList
	resp.Fill(userList)
	resp.Send(w, 200, resp)
}
