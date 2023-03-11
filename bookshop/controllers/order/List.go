package order

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/responses"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger

	user := r.Context().Value("user").(models.User)

	var order models.Order
	order.UserID = user.ID

	orderList, err := order.List()
	if err != nil {
		l.Print("Order_controller", "error", "List", "", "order.List", err.Error())
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
