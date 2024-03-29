CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    PASSWORD TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT(NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT(NOW()),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    public_id uuid UNIQUE DEFAULT gen_random_uuid() NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    owner_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT(NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT(NOW()),
    deleted_at TIMESTAMP,
    CONSTRAINT fk_users FOREIGN KEY(owner_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS project_todo_states (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    item_order INT NOT NULL,
    project_id INT NOT NULL,
    CONSTRAINT fk_project_todo_states_projects FOREIGN KEY (project_id) REFERENCES projects(id)
);

CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    public_id uuid UNIQUE DEFAULT gen_random_uuid() NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    item_order INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT(NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT(NOW()),
    deleted_at TIMESTAMP,
    project_id INT NOT NULL,
    state_id INT NOT NULL,
    CONSTRAINT fk_todo_projects FOREIGN KEY (project_id) REFERENCES projects(id),
    CONSTRAINT fk_todo_state FOREIGN KEY (state_id) REFERENCES project_todo_states(id)
);

CREATE TABLE IF NOT EXISTS todo_assignments (
    todo_id INT NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT fk_todo_user_assignment FOREIGN KEY (todo_id) REFERENCES todos(id),
    CONSTRAINT fk_user_todo_assignment FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS project_assignments (
    project_id INT NOT NULL,
    user_id INT NOT NULL,
    project_owner_id INT NOT NULL,
    CONSTRAINT fk_project_user_assignment FOREIGN KEY (project_id) REFERENCES projects(id),
    CONSTRAINT fk_user_project_assignment FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_project_owner_assignment FOREIGN KEY (project_owner_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS project_invitations (
    project_id INT NOT NULL,
    email TEXT NOT NULL,
    token TEXT NOT NULL,
    sent_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    accepted_at TIMESTAMP,
    CONSTRAINT fk_project_invitation_project FOREIGN KEY (project_id) REFERENCES projects(id)
);
