package apihandlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type registerHandlerAuth interface {
	Register(ctx context.Context, fn string, ln string, em string, pw string) (string, error)
}

func RegisterHandler(as registerHandlerAuth) http.HandlerFunc {
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

		token, err := as.Register(
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
			Domain:   "localhost:3000",
			Name:     "session_token",
			HttpOnly: true,
			Value:    token,
			MaxAge:   86400,
			Path:     "/",
		}
		http.SetCookie(w, &c)

		w.Header().Add("HX-Redirect", "home")
		w.Write([]byte("Success"))
	}
}

type loginHandlerAuth interface {
	Login(ctx context.Context, em string, pw string) (string, error)
}

func LoginHandler(as loginHandlerAuth) http.HandlerFunc {
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

		token, err := as.Login(r.Context(), rb.Email, rb.Password)
		if err != nil {
			fmt.Printf("Can not login: token: %#v, \n error: %s\n", token, err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		c := http.Cookie{
			Domain:   "localhost:3000",
			Name:     "session_token",
			Path:     "/",
			HttpOnly: true,
			Value:    token,
			MaxAge:   86400,
			Expires:  time.Now().Add(86400 * time.Second),
		}
		http.SetCookie(w, &c)

		w.Header().Add("HX-Redirect", "home")
		w.Write([]byte("Success"))
	}
}
