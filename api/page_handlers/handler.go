package pagehandler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/views/pages"
)

// Public Pages
func IndexPage() http.HandlerFunc {
	ch := templ.Handler(pages.Index())
	return func(w http.ResponseWriter, r *http.Request) {
		ch.ServeHTTP(w, r)
	}
}

// Protected Pages
func IndexPageProtected(us *app.UserService, ps *app.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value("user").(app.User)
		initials := string([]byte{u.FirstName[0], u.LastName[0]})

		ch := templ.Handler(pages.IndexProtected(initials))
		ch.ServeHTTP(w, r)
	}
}

func ProjectPageProtected(ps *app.ProjectService, ts *app.ToDoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pubId, err := uuid.Parse(chi.URLParam(r, "projectId"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		p, err := ps.FindProjectById(r.Context(), pubId.String())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		ts, tm, err := ts.FindTodosByProjectId(r.Context(), p.ID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		fmt.Printf("TodoState: %#v\n", ts)
		fmt.Printf("TodoMapColl: %+v\n", tm)
		ch := templ.Handler(pages.SingleProject(ts, tm))
		ch.ServeHTTP(w, r)
	}
}
