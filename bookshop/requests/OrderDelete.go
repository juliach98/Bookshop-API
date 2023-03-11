package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DeleteOrder struct {
	OrderID string `json:"order_id"`
}

func (m *DeleteOrder) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
