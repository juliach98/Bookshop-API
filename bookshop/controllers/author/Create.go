package author

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.CreateAuthor

	err := req.Load(r)

	if err != nil {
		l.Print("Author_controller", "error", "Create", "", "req.Load", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	if len(req.Name) == 0 {
		l.Print("Author_controller", "error", "Create", "", "len(req.Name) == 0", "Name is empty")
		responses.ErrorResponse(w, 502, "Name is empty", "")
		return
	}

	if len(req.Surname) == 0 {
		l.Print("Author_controller", "error", "Create", "", "len(req.Surname) == 0", "Surname is empty")
		responses.ErrorResponse(w, 503, "Surname is empty", "")
		return
	}

	var author models.Author

	author.Name = req.Name
	author.Surname = req.Surname
	author.Bio = req.Bio

	err = author.Create()

	if err != nil {
		l.Print("Author_controller", "error", "Create", "", "author.Create", err.Error())
		responses.ErrorResponse(w, 504, err.Error(), "")
		return
	}

	var resp responses.CreateAuthor
	resp.Fill(author)
	resp.Send(w, 200, resp)
}
