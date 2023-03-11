package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DeleteAuthor struct {
	AuthorID string `json:"author_id"`
}

func (m *DeleteAuthor) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
