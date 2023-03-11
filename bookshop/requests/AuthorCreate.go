package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CreateAuthor struct {
	AuthorID string `json:"author_id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Bio      string `json:"bio"`
}

func (m *CreateAuthor) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
