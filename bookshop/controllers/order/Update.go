package order

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.CreateOrder

	err := req.Load(r)

	if err != nil {
		l.Print("Order_controller", "error", "Update", "", "req.Load", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	if len(req.OrderID) == 0 {
		l.Print("Order_controller", "error", "Delete", "", "len(req.OrderID) == 0", "OrderID is empty")
		responses.ErrorResponse(w, 502, "OrderID is empty", "")
		return
	}

	var order models.Order
	order.StringID = req.OrderID

	found, err := order.Find()

	if err != nil {
		l.Print("Order_controller", "error", "Update", "", "order.Find", err.Error())
		responses.ErrorResponse(w, 503, err.Error(), "")
		return
	}

	if !found {
		l.Print("Order_controller", "error", "Update", "", "order.Find", "Order not found")
		responses.ErrorResponse(w, 404, "Order not found", "")
		return
	}

	if req.DeliveredAt > 0 {
		order.DeliveredAt = req.DeliveredAt
	}

	if req.DeliveryDateTime != order.DeliveryDateTime && !req.DeliveryDateTime.IsZero() {
		order.DeliveryDateTime = req.DeliveryDateTime
	}

	var book models.Book
	var bookOrder models.BookOrder

	booksID := make([]string, len(req.BooksID))

	if len(req.BooksID) > 0 {
		bookList, err := book.ListByStringID(req.BooksID)
		if err != nil {
			l.Print("Order_controller", "error", "Update", "", "book.ListByStringID", err.Error())
			responses.ErrorResponse(w, 504, err.Error(), "")
			return
		}

		for _, book = range bookList {
			bookOrder.BookID = book.ID
			bookOrder.OrderID = order.ID

			err := bookOrder.Create()
			if err != nil {
				l.Print("Order_controller", "error", "Update", "", "bookOrder.Create", err.Error())
				responses.ErrorResponse(w, 505, err.Error(), "")
				return
			}

			booksID = append(booksID, book.StringID)
		}
	}

	err = order.Update()

	if err != nil {
		l.Print("Order_controller", "error", "Update", "", "order.Update", err.Error())
		responses.ErrorResponse(w, 506, err.Error(), "")
		return
	}

	var resp responses.CreateOrder
	resp.Fill(order, booksID)
	resp.Send(w, 200, resp)
}
