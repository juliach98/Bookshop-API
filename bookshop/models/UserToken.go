package models

import (
	"bookshop/drivers"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"strconv"
	"time"
)

type UserToken struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func (m *UserToken) FindBearer() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM users_token WHERE token = $1 LIMIT 1", m.Token)

	err = m.ScanRow(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (m *UserToken) Find() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM users_token WHERE user_id = $1 LIMIT 1", m.UserID)

	err = m.ScanRow(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (m *UserToken) CreateToken(userID int64) error {
	m.UserID = userID

	text := strconv.FormatInt(userID, 10) + strconv.FormatInt(time.Now().UnixNano(), 10)

	hash := md5.Sum([]byte(text))
	m.Token = hex.EncodeToString(hash[:])

	err := m.Create()

	return err
}

func (m *UserToken) Create() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	row := db.QueryRow("INSERT INTO users_token(user_id, token) VALUES ($1,$2) RETURNING ID",
		m.UserID, m.Token)

	err = row.Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil

}

func (m *UserToken) Delete() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("DELETE FROM users_token WHERE user_id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserToken) ScanRow(row *sql.Row) error {
	return row.Scan(
		&m.ID,
		&m.UserID,
		&m.Token,
	)
}
