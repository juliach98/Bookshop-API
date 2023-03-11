package models

import (
	"bookshop/drivers"
	"database/sql"
)

type Address struct {
	ID              int64  `json:"id"`
	Country         string `json:"country"`
	City            string `json:"city"`
	Street          string `json:"street"`
	HouseNumber     string `json:"house_number"`
	ApartmentNumber string `json:"apartment_number"`
	Floor           int64  `json:"floor"`
	UserID          int64  `json:"user_id"`
}

func (m *Address) Create() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("INSERT INTO addresses(country, city, street, house_number, apartment_number, floor, user_id) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING ID",
		m.Country, m.City, m.Street, m.HouseNumber, m.ApartmentNumber, m.Floor, m.UserID)

	err = row.Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Address) Update() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("UPDATE addresses SET country = $1, city = $2, street = $3, house_number = $4, apartment_number = $5, floor = $6, user_id = $7 WHERE id = $8")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.Country, m.City, m.Street, m.HouseNumber, m.ApartmentNumber, m.Floor, m.UserID, m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Address) Delete() error {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	stmt, err := db.Prepare("DELETE FROM addresses WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Address) Find() (bool, error) {
	db, err := drivers.PostgreSQLConnection()

	if err != nil {
		return false, err
	}

	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM addresses WHERE user_id = $1 LIMIT 1", m.UserID)

	err = m.ScanRow(row)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil

}

func (m *Address) ScanRow(rows *sql.Row) error {
	return rows.Scan(
		&m.ID,
		&m.Country,
		&m.City,
		&m.Street,
		&m.HouseNumber,
		&m.ApartmentNumber,
		&m.Floor,
		&m.UserID,
	)
}
