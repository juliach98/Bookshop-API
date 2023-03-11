package responses

import (
	"bookshop/interfaces"
	"bookshop/models"
)

type User struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UserGet struct {
	interfaces.Response
	Data User `json:"data"`
}

func (m *UserGet) Fill(user models.User) {
	m.Status = 200
	m.Error = ""

	m.Data.UserID = user.StringID
	m.Data.Name = user.Name
	m.Data.Surname = user.Surname
	m.Data.Email = user.Email
	m.Data.PhoneNumber = user.PhoneNumber
}
