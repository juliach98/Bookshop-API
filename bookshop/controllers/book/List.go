package book

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/responses"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger

	var book models.Book

	bookList, err := book.List()
	if err != nil {
		l.Print("Book_controller", "error", "List", "", "book.List", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	var bookAuthor models.BookAuthor
	var auth models.Author

	authorsID := make([][]string, len(bookList))

	for i, book := range bookList {
		authorList, err := bookAuthor.List(book.ID, 0)
		if err != nil {
			l.Print("Book_controller", "error", "List", "", "bookAuthor.List", err.Error())
			responses.ErrorResponse(w, 502, err.Error(), "")
			return
		}

		for _, author := range authorList {
			auth.ID = author.AuthorID

			found, err := auth.Find()

			if err != nil {
				l.Print("Book_controller", "error", "List", "", "author.Find", err.Error())
				responses.ErrorResponse(w, 503, err.Error(), "")
				return
			}

			if !found {
				l.Print("Book_controller", "error", "List", "", "author.Find", "Author not found")
				responses.ErrorResponse(w, 404, "Author not found", "")
				return
			}

			authorsID[i] = append(authorsID[i], auth.StringID)
		}
	}

	var resp responses.BookList
	resp.Fill(bookList, authorsID)
	resp.Send(w, 200, resp)
}
