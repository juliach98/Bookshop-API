package users

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/responses"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger

	user := r.Context().Value("user").(models.User)

	var token models.UserToken
	token.UserID = user.ID

	err := token.Delete()

	if err != nil {
		l.Print("User_controller", "error", "Delete", "", "token.Delete", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	var order models.Order
	order.UserID = user.ID

	orderList, err := order.List()

	if err != nil {
		l.Print("User_controller", "error", "Delete", "", "order.List", err.Error())
		responses.ErrorResponse(w, 502, err.Error(), "")
		return
	}

	for _, ord := range orderList {
		var bookOrder models.BookOrder
		bookOrder.OrderID = ord.ID

		err = bookOrder.Delete()

		if err != nil {
			l.Print("User_controller", "error", "Delete", "", "bookOrder.Delete", err.Error())
			responses.ErrorResponse(w, 503, err.Error(), "")
			return
		}

		err = ord.Delete()

		if err != nil {
			l.Print("User_controller", "error", "Delete", "", "ord.Delete", err.Error())
			responses.ErrorResponse(w, 504, err.Error(), "")
			return
		}
	}

	found, err := user.Find()

	if err != nil {
		l.Print("User_controller", "error", "Delete", "", "user.Find", err.Error())
		responses.ErrorResponse(w, 505, err.Error(), "")
		return
	}

	if !found {
		l.Print("User_controller", "error", "Delete", "", "user.Find", "User not found")
		responses.ErrorResponse(w, 404, "User not found", "")
		return
	}

	err = user.Delete()

	if err != nil {
		l.Print("User_controller", "error", "Delete", "", "user.Delete", err.Error())
		responses.ErrorResponse(w, 506, err.Error(), "")
		return
	}

	responses.EmptyResponse(w)
}
