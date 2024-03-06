CREATE TABLE IF NOT EXISTS project_todo_states (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    item_order INT NOT NULL,
    project_id INT NOT NULL,
    CONSTRAINT fk_project_todo_states_projects FOREIGN KEY (project_id) REFERENCES projects(id)
);
