ALTER TABLE
    todos
ADD
    state_id INT NOT NULL;

ALTER TABLE
    todos
ADD
    item_order INT NOT NULL;

ALTER TABLE
    todos
ADD
    CONSTRAINT fk_todo_state FOREIGN KEY (state_id) REFERENCES project_todo_states(id);
