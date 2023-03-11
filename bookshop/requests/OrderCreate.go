package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type CreateOrder struct {
	OrderID          string    `json:"order_id"`
	BooksID          []string  `json:"books_id"`
	DeliveryDateTime time.Time `json:"delivery_date_time"`
	CreatedAt        int64     `json:"created_at"`
	DeliveredAt      int64     `json:"delivered_at"`
}

func (m *CreateOrder) Load(r *http.Request) error {
	rawData, _ := ioutil.ReadAll(r.Body)

	return json.Unmarshal(rawData, &m)
}
