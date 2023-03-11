package book

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/responses"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func SetPicture(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	bookID := r.FormValue("book_id")
	file, handler, err := r.FormFile("picture")
	if err != nil {
		l.Print("Book_controller", "error", "SetPicture", "", "r.FormFile", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}
	defer func() { _ = file.Close() }()

	var book models.Book
	book.StringID = bookID
	found, err := book.FindByStringID()

	if err != nil {
		l.Print("Book_controller", "error", "SetPicture", "", "book.FindByStringID", err.Error())
		responses.ErrorResponse(w, 502, err.Error(), "")
		return
	}

	if !found {
		l.Print("Book_controller", "error", "SetPicture", "", "book.FindByStringID", "Child not found")
		responses.ErrorResponse(w, 404, "Book not found", "")
		return
	}

	bookPicture := bookID + "_" + filepath.Ext(handler.Filename)

	f, err := os.Create("./pictures/books/" + bookPicture)
	if err != nil {
		l.Print("Book_controller", "error", "SetPicture", "", "os.Create", err.Error())
		responses.ErrorResponse(w, 503, err.Error(), "")
		return
	}

	defer func() { _ = f.Close() }()

	if _, err = io.Copy(f, file); err != nil {
		l.Print("Book_controller", "error", "SetPicture", "", "io.Copy", err.Error())
		responses.ErrorResponse(w, 504, err.Error(), "")
		return
	}

	domain, _ := os.LookupEnv("DOMAIN")

	book.Picture = domain + "/pictures/books/" + bookPicture

	err = book.Update()
	if err != nil {
		l.Print("Book_controller", "error", "SetPicture", "", "book.Update", err.Error())
		responses.ErrorResponse(w, 505, err.Error(), "")
		return
	}

	responses.EmptyResponse(w)
}
