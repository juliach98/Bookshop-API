package main

import (
	"bookshop/controllers/address"
	"bookshop/controllers/author"
	"bookshop/controllers/book"
	"bookshop/controllers/order"
	"bookshop/controllers/users"
	"bookshop/helpers"
	"bookshop/middlewares"
	"bookshop/responses"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	rand.Seed(time.Now().UnixNano())
}

func Error404(w http.ResponseWriter, r *http.Request) {
	responses.ErrorResponse(w, 404, "Page not found", "")
}

func NoIndex(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	var l helpers.Logger

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/register", users.Register).Methods("POST")
	api.HandleFunc("/login", users.Login).Methods("POST")

	api.HandleFunc("/books", book.List).Methods("GET")
	api.HandleFunc("/authors", author.List).Methods("GET")

	add := api.PathPrefix("/address").Subrouter()
	add.Use(middlewares.BearerToken)
	add.HandleFunc("/create", address.Create).Methods("POST")
	add.HandleFunc("/update", address.Update).Methods("POST")
	add.HandleFunc("", address.Delete).Methods("DELETE")
	add.HandleFunc("", address.Get).Methods("GET")

	auth := api.PathPrefix("/author").Subrouter()
	auth.Use(middlewares.BearerToken)
	auth.HandleFunc("/create", middlewares.IsAdmin(author.Create)).Methods("POST")
	auth.HandleFunc("/update", middlewares.IsAdmin(author.Update)).Methods("POST")
	auth.HandleFunc("", middlewares.IsAdmin(author.Delete)).Methods("DELETE")

	bk := api.PathPrefix("/book").Subrouter()
	bk.Use(middlewares.BearerToken)
	bk.HandleFunc("/create", middlewares.IsAdmin(book.Create)).Methods("POST")
	bk.HandleFunc("/update", middlewares.IsAdmin(book.Update)).Methods("POST")
	bk.HandleFunc("/picture", middlewares.IsAdmin(book.SetPicture)).Methods("POST")
	bk.HandleFunc("", middlewares.IsAdmin(book.Delete)).Methods("DELETE")

	user := api.PathPrefix("/user").Subrouter()
	user.Use(middlewares.BearerToken)
	user.HandleFunc("/update", users.Update).Methods("POST")
	user.HandleFunc("/logout", users.Logout).Methods("POST")
	user.HandleFunc("", users.Delete).Methods("DELETE")
	user.HandleFunc("", users.Get).Methods("GET")
	user.HandleFunc("/list", middlewares.IsAdmin(users.List)).Methods("GET")

	ord := api.PathPrefix("/order").Subrouter()
	ord.Use(middlewares.BearerToken)
	ord.HandleFunc("/create", order.Create).Methods("POST")
	ord.HandleFunc("/update", order.Update).Methods("POST")
	ord.HandleFunc("/list", order.List).Methods("GET")
	ord.HandleFunc("", order.Delete).Methods("DELETE")
	ord.HandleFunc("/list/{userID:[0-9]+}", middlewares.IsAdmin(order.UserList)).Methods("GET")

	fileRouter := r.PathPrefix("/pictures").Subrouter()
	fileRouter.Use(NoIndex)
	fileRouter.PathPrefix("/").Handler(http.StripPrefix("/pictures", http.FileServer(http.Dir("./pictures"))))

	r.NotFoundHandler = r.NewRoute().HandlerFunc(Error404).GetHandler()

	//cors optionsGoes  Below
	corsConfig := cors.New(cors.Options{
		AllowedHeaders: []string{
			"*",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"DELETE",
		},
	})
	handler := corsConfig.Handler(r)

	errorListen := http.ListenAndServe(":8080", handler)
	if errorListen != nil {
		l.Print("Listen", "", "main.go", "", "ListenAndServe", errorListen.Error())
	}
}
