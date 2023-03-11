package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ListAuthor struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (m *ListAuthor) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
