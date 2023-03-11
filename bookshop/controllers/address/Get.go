package address

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/responses"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger

	user := r.Context().Value("user").(models.User)

	var address models.Address
	address.UserID = user.ID

	found, err := address.Find()

	if err != nil {
		l.Print("Address_controller", "error", "Get", "", "address.Find", err.Error())
		responses.ErrorResponse(w, 502, err.Error(), "")
		return
	}

	if !found {
		l.Print("Address_controller", "error", "Get", "", "address.Find", "Address not found")
		responses.ErrorResponse(w, 404, "Address not found", "")
		return
	}

	var resp responses.CreateAddress
	resp.Fill(address)
	resp.Send(w, 200, resp)
}
