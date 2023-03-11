package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ListBook struct {
	Title     string `json:"title"`
	Series    string `json:"series"`
	Publisher string `json:"publisher"`
	Language  string `json:"language"`
}

func (m *ListBook) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
