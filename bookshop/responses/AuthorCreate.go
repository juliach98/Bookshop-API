package responses

import (
	"bookshop/interfaces"
	"bookshop/models"
)

type CreateAuthor struct {
	interfaces.Response
	Data struct {
		AuthorID string `json:"author_id"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Bio      string `json:"bio"`
	} `json:"data"`
}

func (m *CreateAuthor) Fill(author models.Author) {
	m.Data.AuthorID = author.StringID
	m.Data.Name = author.Name
	m.Data.Surname = author.Surname
	m.Data.Bio = author.Bio
}
