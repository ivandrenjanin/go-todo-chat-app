package pg

import "github.com/ivandrenjanin/go-chat-app/app"

func (u *User) ConvertToUser() app.User {
	return app.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}
}

func (p *Project) ConvertToProject() app.Project {
	return app.Project{
		ID:          p.ID,
		PublicID:    p.PublicID,
		Name:        p.Name,
		Description: p.Description,
		OwnerID:     p.OwnerID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (pa *ProjectAssignment) ConvertToProjectAssignment() app.ProjectAssignment {
	return app.ProjectAssignment{
		ProjectID:      pa.ProjectID,
		UserID:         pa.UserID,
		ProjectOwnerID: pa.ProjectOwnerID,
	}
}
