package api

import (
	"context"
	"net/http"

	"github.com/ivandrenjanin/go-chat-app/app"
)

func MakeMiddleware(as *app.AuthService, us *app.UserService) func(http.Handler) http.HandlerFunc {
	return func(next http.Handler) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("session_token")
			if err != nil {
				http.Error(
					w,
					http.StatusText(http.StatusUnauthorized),
					http.StatusUnauthorized,
				)
				return
			}

			token := c.Value
			claims, ok := as.ValidateToken(token)
			if !ok {
				http.Error(
					w,
					http.StatusText(http.StatusUnauthorized),
					http.StatusUnauthorized,
				)
				return
			}

			u, err := us.FindById(r.Context(), claims.UserID)
			if err != nil {
				http.Error(
					w,
					http.StatusText(http.StatusUnauthorized),
					http.StatusUnauthorized,
				)
				return

			}

			ctx := context.WithValue(r.Context(), "user", u)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
