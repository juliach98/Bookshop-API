package responses

import (
	"bookshop/interfaces"
	"bookshop/models"
)

type Author struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Bio     string `json:"bio"`
}

type AuthorList struct {
	interfaces.Response
	Data struct {
		Authors []Author `json:"authors"`
	} `json:"data"`
}

func (m *AuthorList) Fill(authors []models.Author) {
	m.Status = 200
	m.Error = ""

	for _, author := range authors {
		var a Author

		a.ID = author.StringID
		a.Name = author.Name
		a.Surname = author.Surname
		a.Bio = author.Bio

		m.Data.Authors = append(m.Data.Authors, a)
	}
}
