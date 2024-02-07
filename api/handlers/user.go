package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/ivandrenjanin/go-chat-app/database"
)

func CreateUserHandlers(r *chi.Mux, db database.Database) {
	r.Route("/api/users", func(r chi.Router) {
		r.Delete("/{id}", MakeHandler(deleteUser, db))
		r.Get("/{id}", MakeHandler(getUser, db))
		r.Post("/", MakeHandler(insertUser, db))
	})
}

func MakeHandler(h http.HandlerFunc, db database.Database) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", &db)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetDB(r *http.Request) (*database.Database, error) {
	ctx := r.Context()
	db, ok := ctx.Value("db").(*database.Database)
	if !ok {
		return nil, fmt.Errorf("unable to get db")
	}

	return db, nil
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := GetDB(r)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity,
		)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity,
		)
		return
	}
	err = db.Queries.DeleteUser(r.Context(), int32(id))
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity,
		)
		return
	}
	render.JSON(w, r, struct{ Message string }{Message: "success"})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	db, err := GetDB(r)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity,
		)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity,
		)
		return
	}

	user, err := db.Queries.User(r.Context(), int32(id))
	if err != nil {
		fmt.Println(err.Error())
		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity,
		)
		return
	}

	render.JSON(w, r, user)
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	db, err := GetDB(r)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity,
		)
		return
	}
	var userParams database.InsertUserParams

	if err := render.DecodeJSON(r.Body, &userParams); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = db.Queries.InsertUser(r.Context(), userParams)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity,
		)
		return
	}

	render.JSON(w, r, userParams)
	fmt.Printf("db: %#v\n", userParams)
}
