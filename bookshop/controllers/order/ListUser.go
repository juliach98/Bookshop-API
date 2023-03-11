package order

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/responses"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger

	userIDs := mux.Vars(r)["userID"]
	userID, _ := strconv.ParseInt(userIDs, 10, 64)

	var order models.Order
	order.UserID = userID

	orderList, err := order.List()
	if err != nil {
		l.Print("Order_controller", "error", "UserList", "", "order.List", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	booksID := make([][]string, len(orderList))

	for i, ord := range orderList {
		var orderBook models.BookOrder
		orderBookList, err := orderBook.List(0, ord.ID)
		if err != nil {
			l.Print("Order_controller", "error", "UserList", "", "bookAuthor.List", err.Error())
			responses.ErrorResponse(w, 502, err.Error(), "")
			return
		}

		for _, ordBook := range orderBookList {
			var book models.Book
			book.ID = ordBook.BookID

			found, err := book.Find()

			if err != nil {
				l.Print("Order_controller", "error", "UserList", "", "order.Find", err.Error())
				responses.ErrorResponse(w, 503, err.Error(), "")
				return
			}

			if !found {
				l.Print("Order_controller", "error", "UserList", "", "order.Find", "Order not found")
				responses.ErrorResponse(w, 404, "Order not found", "")
				return
			}

			booksID[i] = append(booksID[i], book.StringID)
		}
	}

	var resp responses.OrderList
	resp.Fill(orderList, booksID)
	resp.Send(w, 200, resp)
}
