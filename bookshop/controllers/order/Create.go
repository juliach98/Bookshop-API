package order

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/requests"
	"bookshop/responses"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var l helpers.Logger
	var req requests.CreateOrder

	err := req.Load(r)

	if err != nil {
		l.Print("Order_controller", "error", "Create", "", "req.Load", err.Error())
		responses.ErrorResponse(w, 501, err.Error(), "")
		return
	}

	if len(req.BooksID) == 0 {
		l.Print("Book_controller", "error", "Create", "", "len(req.BooksID) == 0", "DeliveryDateTime is empty")
		responses.ErrorResponse(w, 502, "BooksID is empty", "")
		return
	}

	if req.DeliveryDateTime.IsZero() {
		l.Print("Book_controller", "error", "Delete", "", "req.DeliveryDateTime.IsZero", "DeliveryDateTime is empty")
		responses.ErrorResponse(w, 503, "DeliveryDateTime is empty", "")
		return
	}

	user := r.Context().Value("user").(models.User)

	var order models.Order

	order.UserID = user.ID
	order.DeliveryDateTime = req.DeliveryDateTime

	err = order.Create()

	if err != nil {
		l.Print("Order_controller", "error", "Create", "", "order.Create", err.Error())
		responses.ErrorResponse(w, 504, err.Error(), "")
		return
	}

	var b models.Book

	bookList, err := b.ListByStringID(req.BooksID)
	if err != nil {
		l.Print("Order_controller", "error", "Create", "", "book.ListByStringID", err.Error())
		responses.ErrorResponse(w, 505, err.Error(), "")
		return
	}

	var bookOrder models.BookOrder
	var bookAuthor models.BookAuthor
	var auth models.Author

	authors := make([][]models.Author, len(bookList))

	for i, book := range bookList {
		bookOrder.BookID = book.ID
		bookOrder.OrderID = order.ID

		err := bookOrder.Create()
		if err != nil {
			l.Print("Order_controller", "error", "Create", "", "bookOrder.Create", err.Error())
			responses.ErrorResponse(w, 506, err.Error(), "")
			return
		}

		authorList, err := bookAuthor.List(book.ID, 0)
		if err != nil {
			l.Print("Order_controller", "error", "c", "", "bookAuthor.List", err.Error())
			responses.ErrorResponse(w, 507, err.Error(), "")
			return
		}

		for _, author := range authorList {
			auth.ID = author.AuthorID

			found, err := auth.Find()

			if err != nil {
				l.Print("Order_controller", "error", "Create", "", "author.Find", err.Error())
				responses.ErrorResponse(w, 508, err.Error(), "")
				return
			}

			if !found {
				l.Print("Order_controller", "error", "Create", "", "author.Find", "Author not found")
				responses.ErrorResponse(w, 404, "Author not found", "")
				return
			}

			authors[i] = append(authors[i], auth)
		}
	}

	var resp responses.CreateOrder
	resp.Fill(order, req.BooksID)
	resp.Send(w, 200, resp)
}
