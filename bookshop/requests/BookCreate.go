package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CreateBook struct {
	BookID      string   `json:"book_id"`
	Title       string   `json:"title"`
	Series      string   `json:"series"`
	Price       float64  `json:"price"`
	AuthorsID   []string `json:"authors_id"`
	Publisher   string   `json:"publisher"`
	Language    string   `json:"language"`
	Description string   `json:"description"`
	Count       int64    `json:"count"`
}

func (m *CreateBook) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
