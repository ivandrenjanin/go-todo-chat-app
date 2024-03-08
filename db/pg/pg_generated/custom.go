package pg

import "github.com/ivandrenjanin/go-chat-app/app"

func (u *User) Convert() app.User {
	return app.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}
}

func (p *Project) Convert() app.Project {
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

func (pa *ProjectAssignment) Convert() app.ProjectAssignment {
	return app.ProjectAssignment{
		ProjectID:      pa.ProjectID,
		UserID:         pa.UserID,
		ProjectOwnerID: pa.ProjectOwnerID,
	}
}

func (pi *ProjectInvitation) Convert() app.ProjectInvitation {
	return app.ProjectInvitation{
		ProjectID: pi.ProjectID,
		Email:     pi.Email,
		SentAt:    pi.SentAt,
		ExpiresAt: pi.ExpiresAt,
	}
}

func (t *Todo) Convert() app.Todo {
	return app.Todo{
		ID:          t.ID,
		StateID:     t.StateID,
		PublicID:    t.PublicID,
		Name:        t.Name,
		Description: t.Description,
		ItemOrder:   t.ItemOrder,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func (ts *ProjectTodoState) Convert() app.TodoState {
	return app.TodoState{
		ID:        ts.ID,
		Name:      ts.Name,
		ItemOrder: ts.ItemOrder,
		ProjectID: ts.ProjectID,
	}
}
