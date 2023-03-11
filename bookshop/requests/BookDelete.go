package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DeleteBook struct {
	BookID string `json:"book_id"`
}

func (m *DeleteBook) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
