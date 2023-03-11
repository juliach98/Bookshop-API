package book

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.CreateBook

	err := req.Load(r)

	if err != nil {
		l.Print("Book_controller", "error", "Update", "", "req.Load", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	if len(req.BookID) == 0 {
		l.Print("Book_controller", "error", "Delete", "", "len(req.BookID) == 0", "BookID is empty")
		responses.ErrorResponse(w, 502, "BookID is empty", "")
		return
	}

	var book models.Book
	book.StringID = req.BookID

	found, err := book.FindByStringID()

	if err != nil {
		l.Print("Book_controller", "error", "Update", "", "book.FindByStringID", err.Error())
		responses.ErrorResponse(w, 503, err.Error(), "")
		return
	}

	if !found {
		l.Print("Book_controller", "error", "Update", "", "book.FindByStringID", "Book not found")
		responses.ErrorResponse(w, 404, "Book not found", "")
		return
	}

	if len(req.Title) > 0 {
		book.Title = req.Title
	}

	if len(req.Series) > 0 {
		book.Series = req.Series
	}

	if req.Price > 0 {
		book.Price = req.Price
	}

	if len(req.Publisher) > 0 {
		book.Publisher = req.Publisher
	}

	if len(req.Language) > 0 {
		book.Language = req.Language
	}

	if len(req.Description) > 0 {
		book.Description = req.Description
	}

	if req.Count > 0 {
		book.Count = req.Count
	}

	var bookAuthor models.BookAuthor
	var auth models.Author

	if len(req.AuthorsID) > 0 {
		authorList, err := auth.ListByStringID(req.AuthorsID)
		if err != nil {
			l.Print("Book_controller", "error", "Update", "", "author.ListByStringID", err.Error())
			responses.ErrorResponse(w, 504, err.Error(), "")
			return
		}

		for _, author := range authorList {
			bookAuthor.BookID = book.ID
			bookAuthor.AuthorID = author.ID

			err := bookAuthor.Create()
			if err != nil {
				l.Print("Book_controller", "error", "Update", "", "bookAuthor.Create", err.Error())
				responses.ErrorResponse(w, 505, err.Error(), "")
				return
			}
		}
	}

	err = book.Update()

	if err != nil {
		l.Print("Book_controller", "error", "Update", "", "book.Update", err.Error())
		responses.ErrorResponse(w, 506, err.Error(), "")
		return
	}

	bookAuthList, err := bookAuthor.List(book.ID, 0)
	if err != nil {
		l.Print("Book_controller", "error", "Update", "", "bookAuthor.List", err.Error())
		responses.ErrorResponse(w, 507, err.Error(), "")
		return
	}

	for _, bookAuth := range bookAuthList {
		auth.ID = bookAuth.AuthorID

		found, err := auth.Find()

		if err != nil {
			l.Print("Author_controller", "error", "Update", "", "author.Find", err.Error())
			responses.ErrorResponse(w, 508, err.Error(), "")
			return
		}

		if !found {
			l.Print("Author_controller", "error", "Update", "", "author.Find", "Author not found")
			responses.ErrorResponse(w, 404, "Author not found", "")
			return
		}
	}

	var resp responses.CreateBook
	resp.Fill(book, req.AuthorsID)
	resp.Send(w, 200, resp)
}
