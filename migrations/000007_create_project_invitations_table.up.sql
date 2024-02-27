CREATE TABLE IF NOT EXISTS project_invitations (
    project_id INT NOT NULL,
    email TEXT NOT NULL,
    sent_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    accepted_at TIMESTAMP,
    CONSTRAINT fk_project_invitation_project FOREIGN KEY (project_id) REFERENCES projects(id)
);
