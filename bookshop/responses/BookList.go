package responses

import (
	"bookshop/interfaces"
	"bookshop/models"
)

type Book struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Series      string   `json:"series"`
	Price       float64  `json:"price"`
	Picture     string   `json:"picture"`
	Publisher   string   `json:"publisher"`
	Language    string   `json:"language"`
	Description string   `json:"description"`
	AuthorsID   []string `json:"authors_id"`
}

type BookList struct {
	interfaces.Response
	Data struct {
		Books []Book `json:"books"`
	} `json:"data"`
}

func (m *BookList) Fill(books []models.Book, authorsID [][]string) {
	m.Status = 200
	m.Error = ""

	for i, book := range books {
		var b Book

		b.ID = book.StringID
		b.Title = book.Title
		b.Series = book.Series
		b.Price = book.Price
		b.Picture = book.Picture
		b.Publisher = book.Publisher
		b.Language = book.Language
		b.Description = book.Description
		b.AuthorsID = authorsID[i]

		m.Data.Books = append(m.Data.Books, b)
	}
}
