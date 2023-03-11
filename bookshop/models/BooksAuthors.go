package models

import (
	"bookshop/drivers"
	"database/sql"
)

type BookAuthor struct {
	ID       int64 `json:"id"`
	BookID   int64 `json:"book_id"`
	AuthorID int64 `json:"author_id"`
}

func (m *BookAuthor) Create() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("INSERT INTO books_authors(book_id, author_id) VALUES($1,$2) RETURNING ID",
		m.BookID, m.AuthorID)

	err = row.Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *BookAuthor) List(bookID int64, authorID int64) ([]BookAuthor, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return nil, err
	}

	defer func() { _ = db.Close() }()

	var query string

	if (bookID > 0) && (authorID > 0) {
		query = "SELECT * FROM books_authors where book_id = $1 and author_id = $2"
	} else {
		query = "SELECT * FROM books_authors where book_id = $1 or author_id = $2"
	}

	rows, err := db.Query(query, bookID, authorID)

	if err != nil {
		return nil, err
	}

	var resp []BookAuthor

	for rows.Next() {
		var ba BookAuthor

		err = ba.ScanRows(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, ba)
	}

	return resp, nil
}

func (m *BookAuthor) ScanRows(rows *sql.Rows) error {
	return rows.Scan(
		&m.ID,
		&m.BookID,
		&m.AuthorID,
	)
}
