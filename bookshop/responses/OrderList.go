package responses

import (
	"bookshop/interfaces"
	"bookshop/models"
	"time"
)

type Order struct {
	ID               string    `json:"id"`
	BooksID          []string  `json:"books_id"`
	DeliveryDateTime time.Time `json:"delivery_date_time"`
	CreatedAt        int64     `json:"created_at"`
	DeliveredAt      int64     `json:"delivered_at"`
}

type OrderList struct {
	interfaces.Response
	Data struct {
		Orders []Order `json:"orders"`
	} `json:"data"`
}

func (m *OrderList) Fill(orders []models.Order, booksID [][]string) {
	m.Status = 200
	m.Error = ""

	for i, order := range orders {
		var o Order

		o.ID = order.StringID
		o.CreatedAt = order.CreatedAt
		o.DeliveryDateTime = order.DeliveryDateTime
		o.CreatedAt = order.CreatedAt
		o.DeliveredAt = order.DeliveredAt
		o.BooksID = booksID[i]

		m.Data.Orders = append(m.Data.Orders, o)
	}
}
