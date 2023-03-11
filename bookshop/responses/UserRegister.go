package responses

import "bookshop/interfaces"

type RegisterUser struct {
	interfaces.Response
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

func (m *RegisterUser) Fill(token string) {
	m.Status = 200
	m.Error = ""

	m.Data.Token = token
}
