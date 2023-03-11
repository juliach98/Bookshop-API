package models

import (
	"bookshop/drivers"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"strconv"
	"time"
)

type Order struct {
	ID               int64     `json:"id"`
	UserID           int64     `json:"user_id"`
	DeliveryDateTime time.Time `json:"delivery_date_time"`
	CreatedAt        int64     `json:"created_at"`
	DeliveredAt      int64     `json:"delivered_at"`
	StringID         string    `json:"string_id"`
}

func (m *Order) CreateStringID() {
	text := strconv.FormatInt(m.UserID, 10) + strconv.FormatInt(m.CreatedAt, 10) + strconv.FormatInt(time.Now().UnixNano(), 10)

	hash := md5.Sum([]byte(text))
	m.StringID = hex.EncodeToString(hash[:])

	return
}

func (m *Order) Create() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	m.CreateStringID()
	m.CreatedAt = time.Now().Unix()

	row := db.QueryRow("INSERT INTO orders(user_id, delivery_date_time, created_at, string_id) VALUES($1,$2,$3,$4) RETURNING ID",
		m.UserID, m.DeliveryDateTime, m.CreatedAt, m.StringID)

	err = row.Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Order) Update() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("UPDATE orders SET delivery_date_time = $1, delivered_at = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.DeliveryDateTime, m.DeliveredAt, m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Order) Delete() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("DELETE FROM orders WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Order) Find() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM orders WHERE string_id = $1 LIMIT 1", m.StringID)

	err = m.ScanRow(row)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil

}

func (m *Order) List() ([]Order, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return nil, err
	}

	defer func() { _ = db.Close() }()

	rows, err := db.Query("SELECT * FROM orders WHERE user_id = $1", m.UserID)

	if err != nil {
		return nil, err
	}

	var resp []Order

	for rows.Next() {
		var order Order

		err = order.ScanRows(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, order)
	}

	return resp, nil
}

func (m *Order) ScanRow(row *sql.Row) error {
	return row.Scan(
		&m.ID,
		&m.UserID,
		&m.DeliveryDateTime,
		&m.CreatedAt,
		&m.DeliveredAt,
		&m.StringID,
	)
}

func (m *Order) ScanRows(rows *sql.Rows) error {
	return rows.Scan(
		&m.ID,
		&m.UserID,
		&m.DeliveryDateTime,
		&m.CreatedAt,
		&m.DeliveredAt,
		&m.StringID,
	)
}
