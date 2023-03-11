package models

import (
	"bookshop/drivers"
	"database/sql"
)

type BookOrder struct {
	ID      int64 `json:"id"`
	BookID  int64 `json:"book_id"`
	OrderID int64 `json:"order_id"`
}

func (m *BookOrder) Create() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("INSERT INTO books_orders(book_id, order_id) VALUES($1,$2) RETURNING ID",
		m.BookID, m.OrderID)

	err = row.Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *BookOrder) List(bookID int64, orderID int64) ([]BookOrder, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return nil, err
	}

	defer func() { _ = db.Close() }()

	var query string

	if (bookID > 0) && (orderID > 0) {
		query = "SELECT * FROM books_orders where book_id = $1 and order_id = $2"
	} else {
		query = "SELECT * FROM books_orders where book_id = $1 or order_id = $2"
	}

	rows, err := db.Query(query, bookID, orderID)

	if err != nil {
		return nil, err
	}

	var resp []BookOrder

	for rows.Next() {
		var bo BookOrder

		err = bo.ScanRows(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, bo)
	}

	return resp, nil
}

func (m *BookOrder) Delete() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("DELETE FROM books_orders WHERE order_id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.OrderID)
	if err != nil {
		return err
	}

	return nil
}

func (m *BookOrder) ScanRows(rows *sql.Rows) error {
	return rows.Scan(
		&m.ID,
		&m.BookID,
		&m.OrderID,
	)
}
