package models

import (
	"bookshop/drivers"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"strconv"
	"time"
)

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	StringID    string `json:"string_id"`
	Password    string `json:"password"`
	IsAdmin     bool   `json:"is_admin"`
}

func (m *User) CreateStringID() {
	text := m.Name + m.Surname + m.Email + m.PhoneNumber + strconv.FormatInt(time.Now().UnixNano(), 10)

	hash := md5.Sum([]byte(text))
	m.StringID = hex.EncodeToString(hash[:])

	return
}

func (m *User) Create() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	m.CreateStringID()

	row := db.QueryRow("INSERT INTO users(name, surname, email, phone_number, string_id, password, is_admin) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING ID",
		m.Name, m.Surname, m.Email, m.PhoneNumber, m.StringID, m.Password, m.IsAdmin)

	err = row.Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *User) Update() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("UPDATE users SET name = $1, surname = $2, email = $3, phone_number = $4, password = $5, is_admin = $6 WHERE id = $7")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.Name, m.Surname, m.Email, m.PhoneNumber, m.Password, m.IsAdmin, m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *User) Delete() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *User) Find() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM users WHERE id = $1", m.ID)

	err = m.ScanRow(row)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil

}

func (m *User) FindByStringID() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM users WHERE string_id = $1", m.StringID)

	err = m.ScanRow(row)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil

}

func (m *User) FindByEmail() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM users WHERE email = $1", m.Email)

	err = m.ScanRow(row)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

func (m *User) Login() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM users WHERE email = $1 and password = $2", m.Email, m.Password)

	err = m.ScanRow(row)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil

}

func (m *User) List() ([]User, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return nil, err
	}

	defer func() { _ = db.Close() }()

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	var resp []User

	for rows.Next() {
		var customer User

		err = customer.ScanRows(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, customer)
	}

	return resp, nil
}

func (m *User) ScanRow(row *sql.Row) error {
	return row.Scan(
		&m.ID,
		&m.Name,
		&m.Surname,
		&m.Email,
		&m.PhoneNumber,
		&m.StringID,
		&m.Password,
		&m.IsAdmin,
	)
}

func (m *User) ScanRows(rows *sql.Rows) error {
	return rows.Scan(
		&m.ID,
		&m.Name,
		&m.Surname,
		&m.Email,
		&m.PhoneNumber,
		&m.StringID,
		&m.Password,
		&m.IsAdmin,
	)
}
