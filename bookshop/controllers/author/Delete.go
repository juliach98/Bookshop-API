package author

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.DeleteAuthor

	err := req.Load(r)

	if err != nil {
		l.Print("Author_controller", "error", "Delete", "", "req.Load", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	if len(req.AuthorID) == 0 {
		l.Print("Author_controller", "error", "Delete", "", "len(req.AuthorID) == 0", "AuthorID is empty")
		responses.ErrorResponse(w, 502, "AuthorID is empty", "")
		return
	}

	var author models.Author
	author.StringID = req.AuthorID

	found, err := author.FindByStringID()

	if err != nil {
		l.Print("Author_controller", "error", "Delete", "", "author.FindByStringID", err.Error())
		responses.ErrorResponse(w, 503, err.Error(), "")
		return
	}

	if !found {
		l.Print("Author_controller", "error", "Delete", "", "author.FindByStringID", "Author not found")
		responses.ErrorResponse(w, 404, "Author not found", "")
		return
	}

	err = author.Delete()

	if err != nil {
		l.Print("Author_controller", "error", "Delete", "", "author.Delete", err.Error())
		responses.ErrorResponse(w, 504, err.Error(), "")
		return
	}

	responses.EmptyResponse(w)
}
