package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m *LoginUser) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
