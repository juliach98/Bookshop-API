package responses

import (
	"bookshop/interfaces"
	"bookshop/models"
)

type CreateOrder struct {
	interfaces.Response
	Data Order `json:"data"`
}

func (m *CreateOrder) Fill(order models.Order, booksID []string) {
	m.Status = 200
	m.Error = ""

	m.Data.ID = order.StringID
	m.Data.DeliveryDateTime = order.DeliveryDateTime
	m.Data.CreatedAt = order.CreatedAt
	m.Data.DeliveredAt = order.DeliveredAt
	m.Data.BooksID = booksID

}
