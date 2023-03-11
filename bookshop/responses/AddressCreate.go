package responses

import (
	"bookshop/interfaces"
	"bookshop/models"
)

type CreateAddress struct {
	interfaces.Response
	Data struct {
		Country         string `json:"country"`
		City            string `json:"city"`
		Street          string `json:"street"`
		HouseNumber     string `json:"house_number"`
		ApartmentNumber string `json:"apartment_number"`
		Floor           int64  `json:"floor"`
	} `json:"data"`
}

func (m *CreateAddress) Fill(address models.Address) {
	m.Status = 200
	m.Error = ""

	m.Data.Country = address.Country
	m.Data.City = address.City
	m.Data.Street = address.Street
	m.Data.HouseNumber = address.HouseNumber
	m.Data.ApartmentNumber = address.ApartmentNumber
	m.Data.Floor = address.Floor
}
