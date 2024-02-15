package apihandlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/views/pages"
)

func RegisterHandler(us *app.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// validate the request body 👍
		// create a new user if it does not exist (email)
		// Write Headers to set an auth cookie ???
		// return IndexProtecte
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
		validate := validator.New()

		fmt.Printf("Req Body: %#v\n", rb)
		err := validate.Struct(rb)
		if err != nil {
			fmt.Printf("There was an error: %s\n", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		userId, err := us.RegisterUser(
			r.Context(),
			rb.FirstName,
			rb.LastName,
			rb.Email,
			rb.Password,
		)
		if err != nil {
			fmt.Printf("There was an error: %s\n", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		fmt.Printf("User Id: %#v\n", userId)
		w.Header().Add("HX-Push-Url", "home")
		ch := templ.Handler(pages.IndexProtected())
		ch.ServeHTTP(w, r)
	}
}

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// validate the request body
		// check if the user exists, if the passwords match
		// Write Headers to set an auth cookie ???
		// return IndexProtected

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading the body %s", err)
			return
		}

		fmt.Printf("Body: %s\n", string(body))
		w.Header().Add("HX-Push-Url", "home")
		ch := templ.Handler(pages.IndexProtected())
		ch.ServeHTTP(w, r)
	}
}
