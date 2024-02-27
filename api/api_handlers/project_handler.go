package apihandlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/views/components"
)

func DeleteProjectHandler(ps *app.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pubId := chi.URLParam(r, "projectId")

		u := r.Context().Value("user").(app.User)
		err := ps.RemoveProject(r.Context(), u, pubId)
		if err != nil {
			msg := err.Error()
			if msg == "Forbidden Operation" {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func renderProjectTableComponent(w http.ResponseWriter, r *http.Request, ps *app.ProjectService) {
	u := r.Context().Value("user").(app.User)
	pc, err := ps.FindProjectsByUserId(r.Context(), u.ID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	headers := []string{"Project Name", "Description", "Actions"}
	rows := make([][]string, 0, cap(pc))
	base := "/api/projects"

	for _, project := range pc {
		subBase := fmt.Sprintf("%s/%s", base, project.PublicID.String())
		assign := fmt.Sprintf("%s/assign", subBase)
		r := []string{assign, subBase, project.Name, project.Description}
		rows = append(rows, r)
	}

	c := templ.Handler(components.ProjectTable(headers, rows))
	c.ServeHTTP(w, r)
}

func GetProjectByUserIdHandler(ps *app.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderProjectTableComponent(w, r, ps)
	}
}

func CreateProjectHandler(ps *app.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type requestBody struct {
			Name        string `validate:"required,min=2,max=32"`
			Description string `validate:"required,min=2,max=32"`
		}

		r.ParseForm()
		var rb requestBody
		rb.Name = r.Form.Get("name")
		rb.Description = r.Form.Get("description")
		if err := validator.New().Struct(rb); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		u := r.Context().Value("user").(app.User)
		_, err := ps.CreateProject(r.Context(), u, rb.Name, rb.Description)
		if err != nil {
			// TODO: Handle this properly
			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
		}

		renderProjectTableComponent(w, r, ps)
	}
}
