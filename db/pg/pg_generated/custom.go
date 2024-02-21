package pg

import "github.com/ivandrenjanin/go-chat-app/app"

func (u User) ConvertToUser() app.User {
	return app.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}
}
