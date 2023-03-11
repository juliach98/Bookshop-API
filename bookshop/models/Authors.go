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

type Author struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Bio      string `json:"bio"`
	StringID string `json:"string_id"`
}

func (m *Author) CreateStringID() {
	text := m.Name + m.Surname + m.Bio + strconv.FormatInt(time.Now().UnixNano(), 10)

	hash := md5.Sum([]byte(text))
	m.StringID = hex.EncodeToString(hash[:])

	return
}

func (m *Author) Create() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	m.CreateStringID()

	row := db.QueryRow("INSERT INTO authors(name, surname, bio, string_id) VALUES($1,$2,$3,$4) RETURNING ID",
		m.Name, m.Surname, m.Bio, m.StringID)

	err = row.Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Author) Update() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("UPDATE authors SET name = $1, surname = $2, bio = $3 WHERE id = $4")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.Name, m.Surname, m.Bio, m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Author) Delete() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("DELETE FROM authors WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Author) Find() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM authors WHERE id = $1 LIMIT 1", m.ID)

	err = m.ScanRow(row)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil

}

func (m *Author) FindByStringID() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM authors WHERE string_id = $1 LIMIT 1", m.StringID)

	err = m.ScanRow(row)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil

}

func (m *Author) List() ([]Author, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return nil, err
	}

	defer func() { _ = db.Close() }()

	rows, err := db.Query("SELECT * FROM authors")

	if err != nil {
		return nil, err
	}

	var resp []Author

	for rows.Next() {
		var author Author

		err = author.ScanRows(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, author)
	}

	return resp, nil
}

func (m *Author) ListByStringID(authorIDs []string) ([]Author, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return nil, err
	}

	defer func() { _ = db.Close() }()

	rows, err := db.Query("SELECT * FROM authors where string_id in ($1)",
		strings.Join(authorIDs, ","))

	if err != nil {
		return nil, err
	}

	var resp []Author

	for rows.Next() {
		var author Author

		err = author.ScanRows(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, author)
	}

	return resp, nil
}

func (m *Author) ScanRow(row *sql.Row) error {
	return row.Scan(
		&m.ID,
		&m.Name,
		&m.Surname,
		&m.Bio,
		&m.StringID,
	)
}

func (m *Author) ScanRows(rows *sql.Rows) error {
	return rows.Scan(
		&m.ID,
		&m.Name,
		&m.Surname,
		&m.Bio,
		&m.StringID,
	)
}
