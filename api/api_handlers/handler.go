package apihandlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/views/pages"
)

func RegisterHandler(us *app.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type requestBody struct {
			FirstName       string `validate:"required,min=2,max=32"`
			LastName        string `validate:"required,min=2,max=32"`
			Email           string `validate:"required,email,min=5,max=32"`
			Password        string `validate:"required,min=8,max=16,eqfield=ConfirmPassword"`
			ConfirmPassword string `validate:"required,min=8,max=16,eqfield=Password"`
		}

		r.ParseForm()
		var rb requestBody
		rb.FirstName = r.Form.Get("first_name")
		rb.LastName = r.Form.Get("last_name")
		rb.Email = r.Form.Get("email")
		rb.Password = r.Form.Get("password")
		rb.ConfirmPassword = r.Form.Get("confirm_password")

		if err := validator.New().Struct(rb); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		token, err := us.RegisterUser(
			r.Context(),
			rb.FirstName,
			rb.LastName,
			rb.Email,
			rb.Password,
		)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		c := http.Cookie{
			Name:     "app-token",
			Value:    token,
			HttpOnly: true,
		}
		http.SetCookie(w, &c)
		w.Header().Add("HX-Push-Url", "home")
		ch := templ.Handler(pages.IndexProtected())
		ch.ServeHTTP(w, r)
	}
}

func LoginHandler(us *app.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type requestBody struct {
			Email    string `validate:"required,email,min=5,max=32"`
			Password string `validate:"required,min=8,max=16"`
		}

		r.ParseForm()
		var rb requestBody
		rb.Email = r.Form.Get("email")
		rb.Password = r.Form.Get("password")

		if err := validator.New().Struct(rb); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		token, err := us.Login(r.Context(), rb.Email, rb.Password)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		c := http.Cookie{
			Name:     "app-token",
			Value:    token,
			HttpOnly: true,
		}
		http.SetCookie(w, &c)
		w.Header().Add("HX-Push-Url", "home")
		ch := templ.Handler(pages.IndexProtected())
		ch.ServeHTTP(w, r)
	}
}
