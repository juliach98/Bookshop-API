package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DeleteAddress struct {
	AddressID string `json:"address_id"`
}

func (m *DeleteAddress) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
