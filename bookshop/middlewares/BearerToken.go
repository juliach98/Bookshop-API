package middlewares

import (
	"bookshop/helpers"
	"bookshop/models"
	"bookshop/responses"
	"context"
	"net/http"
	"strings"
)

func BearerToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		var l helpers.Logger
		var reqToken string

		reqToken = r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")

		if len(splitToken) < 2 {
			responses.NotAuthResponse(writer)
			l.Print("request", "info", r.URL.String(), " ", reqToken, "NotAuthResponse(splitToken)")
			return
		}

		reqToken = splitToken[1]

		var token models.UserToken
		token.Token = reqToken

		if res, err := token.FindBearer(); !res || err != nil {
			responses.NotAuthResponse(writer)
			if err != nil {
				l.Print("middlewares", "error", r.URL.String(), " ", reqToken, err.Error())
			}
			l.Print("request", "info", r.URL.String(), " ", reqToken, "NotAuthResponse(token.FindBearer(reqToken))")
			return
		}

		var user models.User
		user.ID = token.UserID

		res, err := user.Find()

		if !res || err != nil {
			responses.NotAuthResponse(writer)
			if err != nil {
				l.Print("middlewares", "error", r.URL.String(), " ", reqToken, err.Error())
			}

			l.Print("request", "info", r.URL.String(), " ", reqToken, "User model load error")
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		ctx = context.WithValue(ctx, "token", reqToken)
		next.ServeHTTP(writer, r.WithContext(ctx))
	})
}
