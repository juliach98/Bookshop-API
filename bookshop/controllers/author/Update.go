package author

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.CreateAuthor

	err := req.Load(r)

	if err != nil {
		l.Print("Author_controller", "error", "Update", "", "req.Load", err.Error())
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
		l.Print("Author_controller", "error", "Update", "", "author.FindByStringID", err.Error())
		responses.ErrorResponse(w, 503, err.Error(), "")
		return
	}

	if !found {
		l.Print("Author_controller", "error", "Update", "", "author.FindByStringID", "Author not found")
		responses.ErrorResponse(w, 404, "Author not found", "")
		return
	}

	if len(req.Name) > 0 {
		author.Name = req.Name
	}

	if len(req.Surname) > 0 {
		author.Surname = req.Surname
	}

	if len(req.Bio) > 0 {
		author.Bio = req.Bio
	}

	err = author.Update()

	if err != nil {
		l.Print("Author_controller", "error", "Update", "", "author.Update", err.Error())
		responses.ErrorResponse(w, 504, err.Error(), "")
		return
	}

	var resp responses.CreateAuthor
	resp.Fill(author)
	resp.Send(w, 200, resp)
}
