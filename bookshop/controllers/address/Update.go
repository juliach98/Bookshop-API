package address

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.CreateAddress

	err := req.Load(r)

	if err != nil {
		l.Print("Address_controller", "error", "Update", "", "req.Load", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	user := r.Context().Value("user").(models.User)

	var address models.Address
	address.UserID = user.ID

	found, err := address.Find()

	if err != nil {
		l.Print("Address_controller", "error", "Update", "", "address.Find", err.Error())
		responses.ErrorResponse(w, 502, err.Error(), "")
		return
	}

	if !found {
		l.Print("Address_controller", "error", "Update", "", "address.Find", "Address not found")
		responses.ErrorResponse(w, 404, "Address not found", "")
		return
	}

	if len(req.Country) > 0 {
		address.Country = req.Country
	}

	if len(req.City) > 0 {
		address.City = req.City
	}

	if len(req.Street) > 0 {
		address.Street = req.Street
	}

	if len(req.HouseNumber) > 0 {
		address.HouseNumber = req.HouseNumber
	}

	if len(req.ApartmentNumber) > 0 {
		address.ApartmentNumber = req.ApartmentNumber
	}

	address.Floor = req.Floor

	err = address.Update()

	if err != nil {
		l.Print("Address_controller", "error", "Update", "", "address.Update", err.Error())
		responses.ErrorResponse(w, 503, err.Error(), "")
		return
	}

	var resp responses.CreateAddress
	resp.Fill(address)
	resp.Send(w, 200, resp)
}
