package author

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/responses"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger

	var author models.Author

	authorList, err := author.List()
	if err != nil {
		l.Print("Author_controller", "error", "List", "", "author.List", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	var resp responses.AuthorList
	resp.Fill(authorList)
	resp.Send(w, 200, resp)
}
