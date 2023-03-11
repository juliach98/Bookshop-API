package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CreateAddress struct {
	Country         string `json:"country"`
	City            string `json:"city"`
	Street          string `json:"street"`
	HouseNumber     string `json:"house_number"`
	ApartmentNumber string `json:"apartment_number"`
	Floor           int64  `json:"floor"`
}

func (m *CreateAddress) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
