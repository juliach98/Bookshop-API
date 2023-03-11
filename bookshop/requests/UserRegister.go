package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RegisterUser struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (m *RegisterUser) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
