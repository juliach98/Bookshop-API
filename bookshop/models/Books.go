package models

import (
	"bookshop/drivers"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type Book struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Series      string  `json:"series"`
	Price       float64 `json:"price"`
	Picture     string  `json:"picture"`
	Publisher   string  `json:"publisher"`
	Language    string  `json:"language"`
	Description string  `json:"description"`
	Count       int64   `json:"count"`
	StringID    string  `json:"string_id"`
}

func (m *Book) CreateStringID() {
	text := m.Title + m.Series + m.Publisher +
		m.Language + m.Description + strconv.FormatInt(time.Now().UnixNano(), 10)

	hash := md5.Sum([]byte(text))
	m.StringID = hex.EncodeToString(hash[:])

	return
}

func (m *Book) Create() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	m.CreateStringID()

	row := db.QueryRow("INSERT INTO books(title, series, price, picture, publisher, language, description, count, string_id) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING ID",
		m.Title, m.Series, m.Price, m.Picture, m.Publisher, m.Language, m.Description, m.Count, m.StringID)

	err = row.Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Book) Update() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("UPDATE books SET title = $1, series = $2, price = $3, picture = $4, publisher = $5, language = $6, description = $7, count = $8 WHERE id = $9")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.Title, m.Series, m.Price, m.Picture, m.Publisher, m.Language, m.Description, m.Count, m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Book) Delete() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("DELETE FROM books WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Book) FindByStringID() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM books WHERE string_id = $1 LIMIT 1", m.StringID)

	err = m.ScanRow(row)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil

}

func (m *Book) Find() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM books WHERE id = $1 LIMIT 1", m.ID)

	err = m.ScanRow(row)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil

}

func (m *Book) List() ([]Book, error) {
	db, err := drivers.PostgreSQLConnection()

	rows, err := db.Query("SELECT * FROM books")

	if err != nil {
		return nil, err
	}

	var resp []Book

	for rows.Next() {
		var book Book

		err = book.ScanRows(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, book)
	}

	return resp, nil
}

func (m *Book) ListByStringID(bookIDs []string) ([]Book, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return nil, err
	}

	defer func() { _ = db.Close() }()

	rows, err := db.Query("SELECT * FROM books where string_id in ($1)",
		strings.Join(bookIDs, ","))

	if err != nil {
		return nil, err
	}

	var resp []Book

	for rows.Next() {
		var book Book

		err = book.ScanRows(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, book)
	}

	return resp, nil
}

func (m *Book) ScanRow(row *sql.Row) error {
	return row.Scan(
		&m.ID,
		&m.Title,
		&m.Series,
		&m.Price,
		&m.Picture,
		&m.Publisher,
		&m.Language,
		&m.Description,
		&m.Count,
		&m.StringID,
	)
}

func (m *Book) ScanRows(rows *sql.Rows) error {
	return rows.Scan(
		&m.ID,
		&m.Title,
		&m.Series,
		&m.Price,
		&m.Picture,
		&m.Publisher,
		&m.Language,
		&m.Description,
		&m.Count,
		&m.StringID,
	)
}
