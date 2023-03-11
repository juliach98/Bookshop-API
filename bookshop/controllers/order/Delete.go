package order

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.DeleteOrder

	err := req.Load(r)

	if err != nil {
		l.Print("Order_controller", "error", "Delete", "", "req.Load", err.Error())
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
		l.Print("Order_controller", "error", "Delete", "", "order.Find", err.Error())
		responses.ErrorResponse(w, 503, err.Error(), "")
		return
	}

	if !found {
		l.Print("Order_controller", "error", "Delete", "", "order.Find", "Order not found")
		responses.ErrorResponse(w, 404, "Order not found", "")
		return
	}

	var orderBook models.BookOrder

	orderBook.OrderID = order.ID

	err = orderBook.Delete()

	if err != nil {
		l.Print("Order_controller", "error", "Delete", "", "orderBook.Delete", err.Error())
		responses.ErrorResponse(w, 504, err.Error(), "")
		return
	}

	err = order.Delete()

	if err != nil {
		l.Print("Order_controller", "error", "Delete", "", "order.Delete", err.Error())
		responses.ErrorResponse(w, 505, err.Error(), "")
		return
	}

	responses.EmptyResponse(w)
}
