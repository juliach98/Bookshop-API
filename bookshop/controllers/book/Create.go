package book

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.CreateBook

	err := req.Load(r)

	if err != nil {
		l.Print("Book_controller", "error", "Create", "", "req.Load", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	if len(req.Title) == 0 {
		l.Print("Book_controller", "error", "Create", "", "len(req.Title) == 0", "Title is empty")
		responses.ErrorResponse(w, 502, "Title is empty", "")
		return
	}

	if len(req.Language) == 0 {
		l.Print("Book_controller", "error", "Create", "", "len(req.Language) == 0", "Language is empty")
		responses.ErrorResponse(w, 503, "Language is empty", "")
		return
	}

	if len(req.Description) == 0 {
		l.Print("Book_controller", "error", "Create", "", "len(req.Description) == 0", "Description is empty")
		responses.ErrorResponse(w, 504, "Description is empty", "")
		return
	}

	if req.Price == 0 {
		l.Print("Book_controller", "error", "Create", "", "req.Price == 0", "Price is empty")
		responses.ErrorResponse(w, 505, "Price is empty", "")
		return
	}

	var book models.Book

	book.Title = req.Title
	book.Series = req.Series
	book.Price = req.Price
	book.Publisher = req.Publisher
	book.Language = req.Language
	book.Description = req.Description
	book.Count = req.Count

	err = book.Create()

	if err != nil {
		l.Print("Book_controller", "error", "Create", "", "book.Create", err.Error())
		responses.ErrorResponse(w, 506, err.Error(), "")
		return
	}

	var aut models.Author

	authorList, err := aut.ListByStringID(req.AuthorsID)
	if err != nil {
		l.Print("Book_controller", "error", "Create", "", "author.ListByStringID", err.Error())
		responses.ErrorResponse(w, 507, err.Error(), "")
		return
	}

	var bookAuthor models.BookAuthor

	authorsID := make([]string, len(authorList))

	for _, author := range authorList {
		bookAuthor.BookID = book.ID
		bookAuthor.AuthorID = author.ID

		err := bookAuthor.Create()
		if err != nil {
			l.Print("Book_controller", "error", "Create", "", "bookAuthor.Create", err.Error())
			responses.ErrorResponse(w, 508, err.Error(), "")
			return
		}

		authorsID = append(authorsID, author.StringID)
	}

	var resp responses.CreateBook
	resp.Fill(book, authorsID)
	resp.Send(w, 200, resp)
}
