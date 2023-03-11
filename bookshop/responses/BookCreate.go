package responses

import (
	"bookshop/interfaces"
	"bookshop/models"
)

type CreateBook struct {
	interfaces.Response
	Data struct {
		BookID      string   `json:"book_id"`
		Title       string   `json:"title"`
		Series      string   `json:"series"`
		Price       float64  `json:"price"`
		Picture     string   `json:"picture"`
		Publisher   string   `json:"publisher"`
		Language    string   `json:"language"`
		Description string   `json:"description"`
		Count       int64    `json:"count"`
		AuthorsID   []string `json:"authors_id"`
	} `json:"data"`
}

func (m *CreateBook) Fill(book models.Book, authorsID []string) {
	m.Status = 200
	m.Error = ""

	m.Data.BookID = book.StringID
	m.Data.Title = book.Title
	m.Data.Series = book.Series
	m.Data.Price = book.Price
	m.Data.Picture = book.Picture
	m.Data.Publisher = book.Publisher
	m.Data.Language = book.Language
	m.Data.Description = book.Description
	m.Data.Count = book.Count
	m.Data.AuthorsID = authorsID
}
