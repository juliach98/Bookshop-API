package book

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.DeleteBook

	err := req.Load(r)

	if err != nil {
		l.Print("Book_controller", "error", "Delete", "", "req.Load", err.Error())
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
		l.Print("Book_controller", "error", "Delete", "", "book.FindByStringID", err.Error())
		responses.ErrorResponse(w, 503, err.Error(), "")
		return
	}

	if !found {
		l.Print("Book_controller", "error", "Delete", "", "book.FindByStringID", "Book not found")
		responses.ErrorResponse(w, 404, "Book not found", "")
		return
	}

	err = book.Delete()

	if err != nil {
		l.Print("Book_controller", "error", "Delete", "", "book.Delete", err.Error())
		responses.ErrorResponse(w, 504, err.Error(), "")
		return
	}

	responses.EmptyResponse(w)
}
