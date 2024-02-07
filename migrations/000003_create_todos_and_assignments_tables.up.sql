CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    public_id uuid UNIQUE DEFAULT gen_random_uuid() NOT NULL,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(500) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT(NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT(NOW()),
    deleted_at TIMESTAMP,
    project_id INT NOT NULL,
    CONSTRAINT fk_todo_projects FOREIGN KEY (project_id) REFERENCES projects(id)
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
