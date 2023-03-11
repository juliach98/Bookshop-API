package address

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.CreateAddress

	err := req.Load(r)

	if err != nil {
		l.Print("Address_controller", "error", "Create", "", "req.Load", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	if len(req.Country) == 0 {
		l.Print("Address_controller", "error", "Create", "", "len(req.Country) == 0", "Country is empty")
		responses.ErrorResponse(w, 502, "Country is empty", "")
		return
	}

	if len(req.City) == 0 {
		l.Print("Address_controller", "error", "Create", "", "len(req.City) == 0", "City is empty")
		responses.ErrorResponse(w, 503, "City is empty", "")
		return
	}

	if len(req.Street) == 0 {
		l.Print("Address_controller", "error", "Create", "", "len(req.Street) == 0", "Street is empty")
		responses.ErrorResponse(w, 504, "Street is empty", "")
		return
	}

	if len(req.HouseNumber) == 0 {
		l.Print("Address_controller", "error", "Create", "", "len(req.HouseNumber) == 0", "HouseNumber is empty")
		responses.ErrorResponse(w, 505, "HouseNumber is empty", "")
		return
	}

	user := r.Context().Value("user").(models.User)

	var address models.Address

	address.Country = req.Country
	address.City = req.City
	address.Street = req.Street
	address.HouseNumber = req.HouseNumber
	address.ApartmentNumber = req.ApartmentNumber
	address.Floor = req.Floor
	address.UserID = user.ID

	err = address.Create()

	if err != nil {
		l.Print("Address_controller", "error", "Create", "", "address.Create", err.Error())
		responses.ErrorResponse(w, 506, err.Error(), "")
		return
	}

	var resp responses.CreateAddress
	resp.Fill(address)
	resp.Send(w, 200, resp)
}
