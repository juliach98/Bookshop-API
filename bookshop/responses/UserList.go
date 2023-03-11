package responses

import (
	"bookshop/interfaces"
	"bookshop/models"
)

type UserList struct {
	interfaces.Response
	Data struct {
		Users []User `json:"users"`
	} `json:"data"`
}

func (m *UserList) Fill(users []models.User) {
	m.Status = 200
	m.Error = ""

	for _, user := range users {
		var u User

		u.UserID = user.StringID
		u.Name = user.Name
		u.Surname = user.Surname
		u.Email = user.Email
		u.PhoneNumber = user.PhoneNumber

		m.Data.Users = append(m.Data.Users, u)
	}
}
